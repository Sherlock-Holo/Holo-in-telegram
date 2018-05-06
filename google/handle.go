package google

import "github.com/go-telegram-bot-api/telegram-bot-api"

const help = "`/google question`"

func Handle(bot *tgbotapi.BotAPI, message tgbotapi.Message, args string) {
    if args == "" {
        helpReply := tgbotapi.NewMessage(message.Chat.ID, help)
        helpReply.ReplyToMessageID = message.MessageID
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

    bot.Send(reply)
}
