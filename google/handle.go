package google

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
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
		log.Println(err)
		reply = tgbotapi.NewMessage(message.Chat.ID, "error")
		bot.Send(reply)
		return
	}

	answerStr := answer.String()
	if answerStr == "" {
		log.Println("google answer template execute failed")
		return
	}

	reply = tgbotapi.NewMessage(message.Chat.ID, answerStr)

	reply.ReplyToMessageID = message.MessageID
	reply.ParseMode = tgbotapi.ModeHTML

	if _, err := bot.Send(reply); err != nil {
		log.Println(err)
	}
}
