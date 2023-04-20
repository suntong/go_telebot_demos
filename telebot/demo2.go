package main

import (
	"log"
	"os"
	"time"

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

	sayHi := getHiFunc(b)
	b.Handle("/hello", sayHi)
	b.Handle("/hi", sayHi)

	log.Println("Starting bot")
	b.Start()
}

func getHiFunc(b *tb.Bot) func(m *tb.Message) {
	return func(m *tb.Message) {
		log.Println("[bot] Greating command received.")
		if m.Private() {
			b.Send(m.Sender, "hello back to the sender")
		} else {
			b.Send(m.Chat, "hello back "+m.Sender.FirstName+", from the Bot!")
		}
	}
}
