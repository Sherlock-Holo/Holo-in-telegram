package google

import (
	"context"
	"log"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const help = "*/google* `question`"

type Google struct{}

func (g *Google) Handle(msg tgbotapi.Message, ctx context.Context, ch chan<- tgbotapi.Chattable) {
	args := strings.Split(msg.Text, " ")[1:]

	if len(args) == 1 {
		helpReply := tgbotapi.NewMessage(msg.Chat.ID, help)
		helpReply.ReplyToMessageID = msg.MessageID
		helpReply.ParseMode = tgbotapi.ModeMarkdown

		select {
		case <-ctx.Done():
		case ch <- helpReply:
		}
		return
	}

	var reply tgbotapi.MessageConfig

	answer, err := Search(strings.Join(args, " "))
	switch err {
	default:
		log.Println(err)
		reply = tgbotapi.NewMessage(msg.Chat.ID, "bot 犯迷糊了")

	case EmptyResult:
		reply = tgbotapi.NewMessage(msg.Chat.ID, "bot 没有找到结果，并且不是 bot 吃了！！！")

	case nil:
		str := answer.String()
		if str == "" {
			reply = tgbotapi.NewMessage(msg.Chat.ID, "bot 犯迷糊了")
		} else {
			log.Println("arch answer template execute failed")
			reply = tgbotapi.NewMessage(msg.Chat.ID, str)
		}
	}

	reply.ReplyToMessageID = msg.MessageID
	reply.ParseMode = tgbotapi.ModeHTML

	select {
	case <-ctx.Done():
	case ch <- reply:
	}
}
