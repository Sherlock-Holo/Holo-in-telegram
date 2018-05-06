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

    split := strings.Split(text, " ")
    args := strings.Join(split[1:], " ")

    for key, handle := range mux.keyAndHandle {
        if strings.HasPrefix(text, key) {

            handle(mux.Api, message, args)
            return
        }
    }
}

func NewMux(api *tgbotapi.BotAPI) Mux {
    mux := Mux{Api: api}
    mux.keyAndHandle = make(map[string]func(*tgbotapi.BotAPI, tgbotapi.Message, string))

    return mux
}
