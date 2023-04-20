package main

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

var (
	// Universal markup builders.
	menu     = &tele.ReplyMarkup{ResizeKeyboard: true}
	selector = &tele.ReplyMarkup{}

	// Reply buttons.
	btnHelp     = menu.Text("ℹ Help")
	btnSettings = menu.Text("⚙ Settings")

	// Inline buttons.
	//
	// Pressing it will cause the client to
	// send the bot a callback.
	//
	// Make sure Unique stays unique as per button kind
	// since it's required for callback routing to work.
	//
	btnPrev = selector.Data("⬅", "prev")
	btnNext = selector.Data("➡", "next")
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

	menu.Reply(
		menu.Row(btnHelp),
		menu.Row(btnSettings),
	)
	selector.Inline(
		selector.Row(btnPrev, btnNext),
	)

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello!", menu)
	})

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!", selector)
	})

	// On reply button pressed (message)
	b.Handle(&btnHelp, func(c tele.Context) error {
		return c.Send("Here is some help: ...")
	})

	// On inline button pressed (callback)
	b.Handle(&btnPrev, func(c tele.Context) error {
		log.Println("In btnPrev callback")
		return c.Respond()
	})

	b.Start()
}
