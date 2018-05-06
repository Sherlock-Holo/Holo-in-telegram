package arch

import (
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "strings"
)

const help = "*/arch* `package [repo]` , repo: eg: `stable` , `testing` or `core` , `extra`"

func Handle(bot *tgbotapi.BotAPI, message tgbotapi.Message, args string) {
    if args == "" {
        helpReply := tgbotapi.NewMessage(message.Chat.ID, help)
        helpReply.ReplyToMessageID = message.MessageID
        bot.Send(helpReply)
        return
    }

    split := strings.Split(args, " ")

    var (
        answer Answer
        err    error
        reply  tgbotapi.MessageConfig
    )

    switch len(split) {
    case 1:
        answer, err = Query(args, "")

    case 2:
        answer, err = Query(split[0], split[1])
    }

    switch {
    case err == EmptyResult:
        reply = tgbotapi.NewMessage(message.Chat.ID, "no package")

    case err != nil:
        reply = tgbotapi.NewMessage(message.Chat.ID, "error")

    default:
        reply = tgbotapi.NewMessage(message.Chat.ID, answer.String())
    }

    reply.ReplyToMessageID = message.MessageID
    reply.ParseMode = tgbotapi.ModeMarkdown

    bot.Send(reply)
}
