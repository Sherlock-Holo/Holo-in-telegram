use std::fmt;
use std::fmt::Display;
use std::lazy::SyncOnceCell;

use reqwest::{Client, Error, Url};
use serde::export::Formatter;
use serde::Deserialize;

const SEARCH_URL: &str = "https://www.googleapis.com/customsearch/v1";

const UA: &str = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36";

static CLIENT: SyncOnceCell<Client> = SyncOnceCell::new();

#[derive(Deserialize, Debug, Clone)]
struct Item {
    pub title: String,
    pub snippet: String,
    pub link: String,
}

#[derive(Deserialize, Debug, Clone)]
pub struct GoogleResult {
    items: Vec<Item>,
}

impl Display for GoogleResult {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        if self.items.is_empty() {
            return f.write_str("找不到结果啦");
        }

        let item = &self.items[0];

        let answer = format!(
            r#"<strong>Title: </strong>{}
<strong>URL: </strong>{}
"#,
            item.title, item.link
        );

        f.write_str(&answer)
    }
}

pub async fn search(
    question: &str,
    google_key: &str,
    google_cx: &str,
) -> Result<GoogleResult, Error> {
    let client = CLIENT.get_or_init(|| {
        reqwest::Client::builder()
            .gzip(true)
            .user_agent(UA)
            .build()
            .unwrap()
    });

    let mut url: Url = SEARCH_URL.parse().unwrap();

    url.query_pairs_mut().extend_pairs(&[
        ("key", google_key),
        ("cx", google_cx),
        ("num", "1"),
        ("alt", "json"),
        ("q", question),
    ]);

    Ok(client.get(url).send().await?.json().await?)
}
