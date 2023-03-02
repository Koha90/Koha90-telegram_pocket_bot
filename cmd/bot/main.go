package main

import (
	"log"
	"telegram_pocket_bot/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5611890634:AAFvl5C_pXqOjs4LCz2NVyhFT2YMzEu6mGY")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient("106215-d718bdeccac9ca72d5af236")
	if err != nil {
		log.Fatal(err)
	}

	telegramBot := telegram.NewBot(bot, pocketClient, "http://localhost")
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}