package google

import "github.com/go-telegram-bot-api/telegram-bot-api"

func Handle(bot *tgbotapi.BotAPI, message tgbotapi.Message, args string) {
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

    bot.Send(reply)
}
