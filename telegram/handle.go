package telegram

import (
	"context"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Handle interface {
	Handle(msg tgbotapi.Message, ctx context.Context, ch chan<- tgbotapi.Chattable)
}
