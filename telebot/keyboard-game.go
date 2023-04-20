package main

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

var (
	// markup builders.
	selector = &tele.ReplyMarkup{}

	// Reply buttons.
	btnHelp   = selector.Text("ℹ Help")
	btnCancle = selector.Text("⚫️ Cancle")
)

func main() {
	pref := tele.Settings{
		Token:  os.Getenv("TG_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	btnA := selector.Data("1", "A")
	btnB := selector.Data("2", "B")
	btnC := selector.Data("3", "c")
	selector.Reply(
		selector.Row(btnA, btnB, btnC),
		selector.Row(btnHelp, btnCancle),
	)

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Handle("/game", func(c tele.Context) error {
		return c.Send("game!", selector)
	})

	// On reply button pressed (message)
	b.Handle(&btnHelp, func(c tele.Context) error {
		return c.Send("Here is some help: ...")
	})

	// On inline button pressed (callback)
	b.Handle(&btnA, func(c tele.Context) error {
		log.Println("In btnA callback")
		return c.Respond()
	})

	b.Start()
}
