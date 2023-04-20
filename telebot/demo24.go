package main

import (
	"bytes"
	"log"
	"os"
	"time"

	g24 "github.com/suntong/game24"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_BOT_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	playG24 := getG24Func(b)
	b.Handle("/game24", playG24)
	b.Handle("/g24", playG24)

	log.Println("Starting bot")
	b.Start()
}

func getG24Func(b *tb.Bot) func(m *tb.Message) {
	buf := bytes.NewBufferString("")
	g := g24.NewGame(30, buf)
	return func(m *tb.Message) {
		log.Println("[bot] Play command received.")
		buf.Reset()
		g.Play()
		if m.Private() {
			b.Send(m.Sender, buf.String())
		} else {
			b.Send(m.Chat, buf.String())
		}
	}
}
