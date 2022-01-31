package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func run(chatId int64, link string) {
	bot, err := tgbotapi.NewBotAPI(goDotEnvVariable("TOKEN_TLG"))
	if err != nil {
		fmt.Println(err)
	}

	bot.Debug = true
	msg := tgbotapi.NewMessage(chatId, link)
	bot.Send(msg)
}
