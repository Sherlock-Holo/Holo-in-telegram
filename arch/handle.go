package arch

import (
	"context"
	"log"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const help = "*/arch* `package [repo]` , repo: eg: `stable` , `testing`, `aur` or `core` , `extra`"

type Arch struct{}

func (a *Arch) Handle(msg tgbotapi.Message, ctx context.Context, ch chan<- tgbotapi.Chattable) {
	args := strings.Split(msg.Text, " ")[1:]

	if len(args) == 0 {
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

	if len(args) == 2 && strings.ToLower(args[1]) == "aur" {
		answer, err := aurQuery(args[0])

		switch err {
		default:
			log.Println(err)
			reply = tgbotapi.NewMessage(msg.Chat.ID, "哎呀，咱好像把这个 AUR 包吃了")

		case EmptyResult:
			reply = tgbotapi.NewMessage(msg.Chat.ID, "咱并没有找到这个 AUR 包，而且不是咱吃掉了！！！")

		case nil:
			str := answer.String()
			if str != "" {
				reply = tgbotapi.NewMessage(msg.Chat.ID, str)
			} else {
				log.Println("arch answer template execute failed")
				reply = tgbotapi.NewMessage(msg.Chat.ID, "哎呀，咱好像把这个 AUR 包吃了")
			}
		}
	} else {
		answer, err := officialQuery(args[0], args[1:]...)

		switch err {
		default:
			log.Println(err)
			reply = tgbotapi.NewMessage(msg.Chat.ID, "哎呀，咱好像把这个包吃了")

		case nil:
			str := answer.String()
			if str != "" {
				reply = tgbotapi.NewMessage(msg.Chat.ID, str)
			} else {
				log.Println("arch answer template execute failed")
				reply = tgbotapi.NewMessage(msg.Chat.ID, "哎呀，咱好像把这个包吃了")
			}

		case EmptyResult:
			if len(args[1:]) == 0 {
				answer, err := aurQuery(args[0])

				switch err {
				default:
					log.Println(err)
					reply = tgbotapi.NewMessage(msg.Chat.ID, "哎呀，咱好像把这个 AUR 包吃了")

				case EmptyResult:
					reply = tgbotapi.NewMessage(msg.Chat.ID, "咱并没有找到这个 AUR 包，而且不是咱吃掉了！！！")

				case nil:
					str := answer.String()
					if str != "" {
						reply = tgbotapi.NewMessage(msg.Chat.ID, str)
					} else {
						log.Println("arch answer template execute failed")
						reply = tgbotapi.NewMessage(msg.Chat.ID, "哎呀，咱好像把这个 AUR 包吃了")
					}
				}
			} else {
				reply = tgbotapi.NewMessage(msg.Chat.ID, "咱并没有找到这个包，而且不是咱吃掉了！！！")
			}
		}
	}

	reply.ReplyToMessageID = msg.MessageID
	reply.ParseMode = tgbotapi.ModeHTML

	select {
	case <-ctx.Done():
	case ch <- reply:
	}
}
