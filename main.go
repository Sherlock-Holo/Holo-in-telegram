package main

import (
	"flag"
	"log"
	"os"

	"github.com/Sherlock-Holo/Holo-in-telegram/arch"
	"github.com/Sherlock-Holo/Holo-in-telegram/google"
	"github.com/Sherlock-Holo/Holo-in-telegram/telegram"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	token = flag.String("token", "", "bot token")
	debug = flag.Bool("debug", false, "debug mode")
	key   = flag.String("key", "", "google api key")
	cx    = flag.String("cx", "", "google api cx")
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

	google.Key = *key
	google.Cx = *cx

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
