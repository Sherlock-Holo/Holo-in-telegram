#![feature(once_cell)]

use std::env;
use teloxide::Bot;

pub mod arch;
pub mod command;
pub mod google;
pub mod handle;

pub async fn run() -> anyhow::Result<()> {
    teloxide::enable_logging!();

    let bot = Bot::builder().token(env::var("HOLO_BOT_TOKEN")?).build();

    let bot_name = "SherlockHolo_bot";

    teloxide::commands_repl(bot, bot_name, handle::handle).await;

    Ok(())
}
