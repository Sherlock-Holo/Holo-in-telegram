package telegram

import (
	"context"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Mux struct {
	bot   *tgbotapi.BotAPI
	m     map[string]Handle
	ch    chan tgbotapi.Chattable
	ctx   context.Context
	close context.CancelFunc
}

func (m *Mux) Register(key string, handle Handle) {
	if !strings.HasPrefix(key, "/") {
		key = "/" + key
	}
	m.m[key+"@"+m.bot.Self.UserName] = handle
	m.m[key] = handle
}

func (m *Mux) Do(msg tgbotapi.Message) {
	key := strings.Split(msg.Text, " ")[0]

	if handle, ok := m.m[key]; ok {
		handle.Handle(msg, m.ctx, m.ch)
	}
}

func NewMux(api *tgbotapi.BotAPI) Mux {
	mux := Mux{bot: api}
	mux.m = make(map[string]Handle)
	mux.ch = make(chan tgbotapi.Chattable, 30)
	mux.ctx, mux.close = context.WithCancel(context.Background())

	return mux
}
