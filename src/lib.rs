#![feature(once_cell)]

use std::{env, io};

use teloxide::dispatching::repls::CommandReplExt;
use teloxide::Bot;
use tracing::level_filters::LevelFilter;
use tracing::subscriber;
use tracing_log::LogTracer;
use tracing_subscriber::filter::Targets;
use tracing_subscriber::layer::SubscriberExt;
use tracing_subscriber::{fmt, Registry};

use crate::command::Command;

pub mod arch;
pub mod command;
pub mod google;
pub mod handle;

pub async fn run() -> anyhow::Result<()> {
    init_log();

    let bot = Bot::new(env::var("HOLO_BOT_TOKEN")?);

    Command::repl(bot, handle::handle).await;

    Ok(())
}

fn init_log() {
    LogTracer::init().unwrap();

    let layer = fmt::layer()
        .pretty()
        .with_target(true)
        .with_writer(io::stderr);

    let targets = Targets::new()
        .with_target("h2", LevelFilter::OFF)
        .with_default(LevelFilter::DEBUG);

    let layered = Registry::default()
        .with(targets)
        .with(layer)
        .with(LevelFilter::INFO);

    subscriber::set_global_default(layered).unwrap();
}
