use teloxide::utils::command::{BotCommand, ParseError};

#[derive(BotCommand, PartialEq, Debug, Eq, Clone)]
#[command(rename = "lowercase", parse_with = "split")]
pub enum Command {
    #[command(
        description = "*/arch* `package [repo]`, repo: eg: `stable` , `testing`, `aur` or `core` , `extra`",
        parse_with = "parse_arch"
    )]
    Arch(String, Option<Vec<String>>),

    #[command(description = "*/google* `question`")]
    Google(String),
}

fn parse_arch(input: String) -> Result<(String, Option<Vec<String>>), ParseError> {
    if input.is_empty() {
        return Err(ParseError::TooFewArguments {
            expected: 1,
            found: 0,
            message: "at least 1".to_string(),
        });
    }

    let mut cmd_with_args = input.split_whitespace();

    let cmd = cmd_with_args.next().unwrap().to_string();

    let args = cmd_with_args.map(String::from).collect::<Vec<_>>();

    if args.is_empty() {
        Ok((cmd, None))
    } else {
        Ok((cmd, Some(args)))
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn parse_arch() {
        let cmd: Command = Command::parse("/arch linux", "").unwrap();

        assert_eq!(cmd, Command::Arch("linux".to_string(), None));
    }

    #[test]
    fn parse_arch_with_repo() {
        let cmd: Command = Command::parse("/arch linux testing", "").unwrap();

        assert_eq!(
            cmd,
            Command::Arch("linux".to_string(), Some(vec!["testing".to_string()]))
        );
    }

    #[test]
    fn parse_google() {
        let cmd: Command = Command::parse("/google linux", "").unwrap();

        assert_eq!(cmd, Command::Google("linux".to_string()));
    }
}
