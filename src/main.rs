use holo_in_telegram::run;

#[tokio::main(flavor = "current_thread")]
async fn main() -> anyhow::Result<()> {
    run().await
}
