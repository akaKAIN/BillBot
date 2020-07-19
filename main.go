package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"os"
	"strconv"
)

var DATA Data

func Listen(port string) {
	url := fmt.Sprintf(":%v", port)
	log.Fatal(http.ListenAndServe(url, nil))
}

func Send(text string) {
	_, err := DATA.Bot.Send(tgbotapi.NewMessage(int64(DATA.CHAT_ID), text))
	if err != nil {
		log.Printf("Error SENDing: %v", err)
	}
}

func init() {
	DATA.TOKEN = os.Getenv("TOKEN")
	DATA.PORT = os.Getenv("PORT")
	go Listen(DATA.PORT)
	chatID, err := strconv.Atoi(os.Getenv("CHAT_ID"))
	if err != nil {
		log.Fatalln("Error env:", err)
	} else {
		DATA.CHAT_ID = chatID
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(DATA.TOKEN)
	if err != nil {
		log.Fatal("Error bot creation", err)
	}
	log.Println("Access bot creation.")

	if _, err := bot.SetWebhook(tgbotapi.NewWebhook(webhook)); err != nil {
		log.Fatalf("Setting webhook %v", webhook)
	}
	DATA.Bot = bot
	fmt.Printf("%+v", DATA)
	updates := bot.ListenForWebhook("/")
	for update := range updates {
		if update.Message.Text == "баланс" {
			PrintBill()
			continue
		}
		go ParserPipeLine(update.Message.Text)
		text := fmt.Sprintf("%v\n", update.Message.Text)
		go Send(text)
	}
}
