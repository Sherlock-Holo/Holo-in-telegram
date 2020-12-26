use std::env;
use std::lazy::SyncOnceCell;

use teloxide::prelude::UpdateWithCx;
use teloxide::requests::{Request, ResponseResult};
use teloxide::types::{Message, ParseMode};
use teloxide::RequestError;

use crate::arch;
use crate::command::Command;
use crate::google;

static GOOGLE_KEY_AND_CX: SyncOnceCell<Option<(String, String)>> = SyncOnceCell::new();

pub async fn handle(cx: UpdateWithCx<Message>, cmd: Command) -> ResponseResult<()> {
    let answer = match cmd {
        Command::Google(question) => {
            let google_key_and_cx = GOOGLE_KEY_AND_CX.get_or_init(|| {
                match (
                    env::var("HOLO_BOT_GOOGLE_KEY"),
                    env::var("HOLO_BOT_GOOGLE_CX"),
                ) {
                    (Ok(google_key), Ok(google_cx)) => Some((google_key, google_cx)),

                    _ => None,
                }
            });

            if let Some((google_key, google_cx)) = google_key_and_cx {
                google::search(&question, google_key, google_cx)
                    .await
                    .map_err(RequestError::NetworkError)?
                    .to_string()
            } else {
                return Ok(());
            }
        }

        Command::Arch(name, repos) => {
            let repos = match repos {
                None => vec![],
                Some(repos) => repos
                    .into_iter()
                    .filter(|repo| !repo.is_empty())
                    .collect::<Vec<String>>(),
            };

            let repos = repos.iter().map(|repo| repo.as_str()).collect::<Vec<_>>();

            if repos.contains(&"aur") || repos.contains(&"AUR") {
                arch::aur_query(&name)
                    .await
                    .map_err(RequestError::NetworkError)?
                    .to_string()
            } else {
                let result = arch::official_query(&name, &repos)
                    .await
                    .map_err(RequestError::NetworkError)?;

                if !result.is_empty() {
                    result.to_string()
                } else {
                    let result = arch::aur_query(&name)
                        .await
                        .map_err(RequestError::NetworkError)?;

                    if result.is_empty() {
                        "咱没有找到结果，并且不是咱吃了！！！".to_string()
                    } else {
                        result.to_string()
                    }
                }
            }
        }
    };

    let message_id = cx.update.id;

    cx.reply_to(answer)
        .parse_mode(ParseMode::HTML)
        .reply_to_message_id(message_id)
        .send()
        .await?;

    Ok(())
}
