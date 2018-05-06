package google

import (
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "log"
)

const help = "*/google* `question`"

func Handle(bot *tgbotapi.BotAPI, message tgbotapi.Message, args string) {
    if args == "" {
        helpReply := tgbotapi.NewMessage(message.Chat.ID, help)
        helpReply.ReplyToMessageID = message.MessageID
        helpReply.ParseMode = tgbotapi.ModeMarkdown
        bot.Send(helpReply)
        return
    }

    answer, err := Search(args)

    var (
        reply tgbotapi.MessageConfig
    )

    if err != nil {
        reply = tgbotapi.NewMessage(message.Chat.ID, "error")
        bot.Send(reply)
        return
    }

    reply = tgbotapi.NewMessage(message.Chat.ID, answer.String())

    reply.ReplyToMessageID = message.MessageID
    reply.ParseMode = tgbotapi.ModeMarkdown

    if _, err := bot.Send(reply); err != nil {
        log.Println(err)
    }
}
