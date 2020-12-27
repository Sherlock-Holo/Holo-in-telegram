use std::cmp::Ordering;
use std::fmt;
use std::fmt::Display;
use std::lazy::SyncOnceCell;

use chrono::{DateTime, FixedOffset, NaiveDateTime};
use difflib::sequencematcher::SequenceMatcher;
use reqwest::{Client, Error, Method, Url};
use serde::Deserialize;
use serde::export::Formatter;

const OFFICIAL_URL: &str = "https://www.archlinux.org/packages/search/json";
const AUR_URL: &str = "https://aur.archlinux.org/rpc/";

const STABLE_REPOS: [&str; 4] = ["Core", "Extra", "Community", "Multilib"];
const TESTING_REPOS: [&str; 3] = ["Testing", "Community-Testing", "Multilib-Testing"];

static CLIENT: SyncOnceCell<Client> = SyncOnceCell::new();

#[derive(Deserialize, Debug, Clone)]
pub struct OfficialResultInfo {
    #[serde(rename = "pkgname")]
    name: String,

    #[serde(rename = "pkgdesc")]
    desc: String,

    #[serde(rename = "licenses")]
    licenses: Vec<String>,

    #[serde(rename = "pkgver")]
    version: String,

    #[serde(rename = "pkgrel")]
    rel: String,

    #[serde(rename = "arch")]
    arch: String,

    #[serde(rename = "repo")]
    repo: String,
}

#[derive(Deserialize, Debug, Clone)]
pub struct OfficialResult {
    results: Vec<OfficialResultInfo>,
}

impl Display for OfficialResult {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        if self.results.is_empty() {
            return f.write_str("哎呀，咱好像把这个包吃了~");
        }

        let result = &self.results[0];

        let url = format!(
            "https://www.archlinux.org/packages/{}/{}/{}",
            result.repo, result.arch, result.name
        );

        let answer = format!(
            r#"<strong>name: </strong>{}
<strong>description: </strong>{}
<strong>version: </strong>{}
<strong>pkgrel: </strong>{}
<strong>repo: </strong>{}
<strong>url: </strong>{}
"#,
            result.name, result.desc, result.version, result.rel, result.repo, url
        );

        f.write_str(&answer)
    }
}

impl OfficialResult {
    pub fn optimize_result(&mut self, name: &str) {
        self.results.sort_by(|first, second| {
            let first_ratio = SequenceMatcher::new(name, &first.name).ratio();
            let second_ratio = SequenceMatcher::new(name, &second.name).ratio();

            match first_ratio
                .partial_cmp(&second_ratio)
                .unwrap_or(Ordering::Equal)
            {
                Ordering::Greater => Ordering::Less,
                Ordering::Less => Ordering::Greater,
                Ordering::Equal => Ordering::Equal,
            }
        })
    }

    pub fn is_empty(&self) -> bool {
        self.results.is_empty()
    }
}

pub async fn official_query(name: &str, repos: &[&str]) -> Result<OfficialResult, Error> {
    let client = CLIENT.get_or_init(|| reqwest::Client::builder().gzip(true).build().unwrap());

    let mut url: Url = OFFICIAL_URL.parse().unwrap();

    let repos = match repos.len() {
        0 => &[],
        1 => match repos[0] {
            "stable" => &STABLE_REPOS[..],
            "testing" => &TESTING_REPOS[..],

            _ => repos,
        },

        _ => repos,
    };

    url.query_pairs_mut()
        .append_pair("name", name)
        .extend_pairs(repos.iter().map(|repo| ("repo", repo)));

    let mut result = client
        .request(Method::GET, url)
        .send()
        .await?
        .json::<OfficialResult>()
        .await?;

    result.optimize_result(name);

    Ok(result)
}

#[derive(Deserialize, Debug, Clone)]
pub struct AurResultInfo {
    #[serde(rename = "Name")]
    name: String,

    #[serde(rename = "Description")]
    desc: String,

    #[serde(rename = "Version")]
    version: String,

    #[serde(rename = "URL")]
    url: Option<String>,

    #[serde(skip)]
    rel: u32,

    #[serde(rename = "NumVotes")]
    vote: u32,

    #[serde(rename = "OutOfDate")]
    out_of_date: Option<i64>,

    #[serde(rename = "Maintainer")]
    maintainer: Option<String>,
}

#[derive(Deserialize, Debug, Clone)]
pub struct AurResult {
    #[serde(rename = "resultcount")]
    count: usize,

    results: Vec<AurResultInfo>,
}

impl Display for AurResult {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        if self.results.is_empty() {
            return f.write_str("哎呀，咱好像把这个 AUR 包吃了");
        }

        let result = &self.results[0];

        let pkgrel = match result.version.split('-').last() {
            None => 1,
            Some(pkgrel) => pkgrel.parse().unwrap_or(1),
        };

        let aur_url = "https://aur.archlinux.org/packages/".to_string() + &result.name;

        let cst = FixedOffset::east(8 * 3600);

        let out_of_date = result
            .out_of_date
            .map(|date| {
                DateTime::<FixedOffset>::from_utc(NaiveDateTime::from_timestamp(date, 0), cst)
                    .format("%F")
                    .to_string()
            })
            .unwrap_or_else(|| "no".to_string());

        let answer = format!(
            r#"<strong>name: </strong>{}
<strong>description: </strong>{}
<strong>version: </strong>{}
<strong>pkgrel: </strong>{}
<strong>repo: </strong>{}
<strong>url: </strong>{}
<strong>vote: </strong>{}
<strong>out-of-date: </strong>{}
<strong>maintainer: </strong>{}
<strong>AUR: </strong>{}
"#,
            result.name,
            result.desc,
            result.version,
            pkgrel,
            "AUR",
            result.url.as_deref().unwrap_or(""),
            result.vote,
            out_of_date,
            result.maintainer.as_deref().unwrap_or(""),
            aur_url
        );

        f.write_str(&answer)
    }
}

impl AurResult {
    pub fn optimize_result(&mut self, name: &str) {
        self.results.sort_by(|first, second| {
            let first_ratio = SequenceMatcher::new(name, &first.name).ratio();
            let second_ratio = SequenceMatcher::new(name, &second.name).ratio();

            match first_ratio
                .partial_cmp(&second_ratio)
                .unwrap_or(Ordering::Equal)
            {
                Ordering::Greater => Ordering::Less,
                Ordering::Less => Ordering::Greater,
                Ordering::Equal => Ordering::Equal,
            }
        })
    }

    pub fn is_empty(&self) -> bool {
        self.results.is_empty()
    }
}

pub async fn aur_query(name: &str) -> Result<AurResult, Error> {
    let client = CLIENT.get_or_init(|| reqwest::Client::builder().gzip(true).build().unwrap());

    let mut url: Url = AUR_URL.parse().unwrap();

    url.query_pairs_mut()
        .extend_pairs(&[("v", "5"), ("type", "search"), ("arg", name)]);

    let mut result = client
        .request(Method::GET, url)
        .send()
        .await?
        .json::<AurResult>()
        .await?;

    result.optimize_result(name);

    Ok(result)
}
