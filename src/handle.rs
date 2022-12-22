use std::env;
use std::sync::LazyLock;

use teloxide::payloads::SendMessageSetters;
use teloxide::requests::{Requester, ResponseResult};
use teloxide::types::{Message, ParseMode};
use teloxide::Bot;

use crate::arch;
use crate::command::Command;
use crate::google;

static GOOGLE_KEY_AND_CX: LazyLock<Option<(String, String)>> = LazyLock::new(|| {
    match (
        env::var("HOLO_BOT_GOOGLE_KEY"),
        env::var("HOLO_BOT_GOOGLE_CX"),
    ) {
        (Ok(google_key), Ok(google_cx)) => Some((google_key, google_cx)),

        _ => None,
    }
});

struct Answer {
    result: String,
    mode: ParseMode,
}

pub async fn handle(bot: Bot, msg: Message, cmd: Command) -> ResponseResult<()> {
    let answer = match cmd {
        Command::Google(question) => {
            if let Some((google_key, google_cx)) = GOOGLE_KEY_AND_CX.as_ref() {
                if question.is_empty() {
                    Answer {
                        result: "*/google* `question`".to_string(),
                        mode: ParseMode::MarkdownV2,
                    }
                } else {
                    let answer = google::search(&question, google_key, google_cx)
                        .await
                        .map(|result| result.to_string())
                        .unwrap_or_else(|_| "哎呀咱好像在 Google 数据中心中迷路了~".to_string());

                    Answer {
                        result: answer,
                        mode: ParseMode::Html,
                    }
                }
            } else {
                return Ok(());
            }
        }

        Command::Arch(name, repos) => {
            match name {
                None => Answer {
                    result: "*/arch* `package [repo]`, repo: eg: `stable` , `testing`, `aur` or `core` , `extra`".to_string(),
                    mode: ParseMode::MarkdownV2,
                },
                Some(name) => {
                    let repos = match repos {
                        None => vec![],
                        Some(repos) => repos
                            .into_iter()
                            .filter(|repo| !repo.is_empty())
                            .collect::<Vec<String>>(),
                    };

                    let repos = repos.iter().map(|repo| repo.as_str()).collect::<Vec<_>>();

                    let answer = if repos.contains(&"aur") || repos.contains(&"AUR") {
                        arch::aur_query(&name)
                            .await
                            .map(|result| result.to_string())
                            .unwrap_or_else(|err| {
                                log::error!("{}", err);

                                "哎呀咱好像在 AUR 数据库中迷路了~".to_string()
                            })
                    } else if let Ok(result) = arch::official_query(&name, &repos).await {
                        if !result.is_empty() {
                            result.to_string()
                        } else {
                            arch::aur_query(&name)
                                .await
                                .map(|result| {
                                    if result.is_empty() {
                                        "咱没有找到结果，并且不是咱吃了！！！".to_string()
                                    } else {
                                        result.to_string()
                                    }
                                })
                                .unwrap_or_else(|err| {
                                    log::error!("{}", err);

                                    "哎呀咱好像在 AUR 数据库中迷路了~".to_string()
                                })
                        }
                    } else {
                        "哎呀咱好像在 Archlinux 数据库中迷路了~".to_string()
                    };

                    Answer {
                        result: answer,
                        mode: ParseMode::Html,
                    }
                }
            }
        }
    };

    bot.send_message(msg.chat.id, answer.result)
        .parse_mode(answer.mode)
        .reply_to_message_id(msg.id)
        .await?;

    Ok(())
}
