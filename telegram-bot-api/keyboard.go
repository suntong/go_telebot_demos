package main

import (
	"log"
        "fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("1.com","http://1.com"),
		tgbotapi.NewInlineKeyboardButtonSwitch("2sw","open 2"),
		tgbotapi.NewInlineKeyboardButtonData("3","3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4","4"),
		tgbotapi.NewInlineKeyboardButtonData("5","5"),
		tgbotapi.NewInlineKeyboardButtonData("6","6"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	fmt.Print(".")
	for update := range updates {
		if update.CallbackQuery != nil{
			fmt.Print(update)

			bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID,update.CallbackQuery.Data))

			bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID,update.CallbackQuery.Data))
		}
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch update.Message.Text {
			case "open":
				msg.ReplyMarkup = numericKeyboard

			}

			bot.Send(msg)
		}
	}
}