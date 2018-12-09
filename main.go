package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

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

	bot.Debug = *debug

	google.Key = *key
	google.Cx = *cx

	mux := telegram.NewMux(bot)
	mux.Register("arch", new(arch.Arch))
	if google.Key != "" && google.Cx != "" {
		mux.Register("google", new(google.Google))
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	mux.Run()

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt)

	go func() {
		<-signalCh
		bot.StopReceivingUpdates()
		_ = mux.Close()
	}()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		go mux.Do(*update.Message)
	}
}
