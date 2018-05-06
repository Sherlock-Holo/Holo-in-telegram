package main

import (
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "log"
    "github.com/Sherlock-Holo/Holo-in-telegram/telegram"
    "github.com/Sherlock-Holo/Holo-in-telegram/google"
    "github.com/Sherlock-Holo/Holo-in-telegram/arch"
    "os"
    "flag"
)

var (
    token = flag.String("token", "", "bot token")
    debug = flag.Bool("debug", false, "debug mode")
)

func main() {
    flag.Parse()

    if flag.NFlag() == 0 {
        flag.Usage()
        os.Exit(2)
    }

    bot, err := tgbotapi.NewBotAPI(*token)

    if err != nil {
        log.Fatal(err)
    }

    mux := telegram.NewMux(bot)
    mux.Add("google", google.Handle)
    mux.Add("arch", arch.Handle)

    bot.Debug = *debug

    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates, err := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil {
            continue
        }

        go mux.Do(*update.Message)
    }
}
