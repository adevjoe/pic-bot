package main

import (
	"github.com/adevjoe/pic-bot/pkg/handle"
	"github.com/adevjoe/pic-bot/pkg/utils"
	"log"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	utils.InitFlag()

	b, err := tb.NewBot(tb.Settings{
		Token:  utils.BotToken,
		Poller: &tb.LongPoller{Timeout: utils.TickTime},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	// commands
	b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		handle.Start(b, m)
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		handle.Text(b, m)
	})
	b.Handle(tb.OnPhoto, func(m *tb.Message) {
		handle.Photo(b, m)
	})

	b.Start()
}
