use holo_in_telegram::run;

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    run().await
}
