package telegram

import (
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "strings"
)

type Mux struct {
    Api          *tgbotapi.BotAPI
    keyAndHandle map[string]func(*tgbotapi.BotAPI, tgbotapi.Message, string)
}

func (mux *Mux) Add(key string, handle func(*tgbotapi.BotAPI, tgbotapi.Message, string)) {
    mux.keyAndHandle["/"+key] = handle
}

func (mux *Mux) Do(message tgbotapi.Message) {
    text := message.Text

    for key, handle := range mux.keyAndHandle {
        if strings.HasPrefix(text, key) {
            split := strings.Split(text, " ")

            args := strings.Join(split[1:], " ")
            handle(mux.Api, message, args)
            return
        }
    }
}
