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

	b.Handle("/hello", func(m *tb.Message) {
		log.Println("[bot] /hello command received.")
		if m.Private() {
			b.Send(m.Sender, "hello back to the sender")
		} else {
			b.Send(m.Chat, "hello back "+m.Sender.FirstName+", from the Bot!")
		}
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		// all the text messages that weren't
		// captured by existing handlers
		b.Send(m.Chat, "Got a text message from "+
			m.Sender.FirstName+" "+m.Sender.LastName+": "+m.Text)
		//log.Println(m.Payload)
	})

	log.Println("Starting bot")
	b.Start()
}
