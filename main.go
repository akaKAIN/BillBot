package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Listen(port string) {
	url := fmt.Sprintf(":%v", port)
	log.Fatal(http.ListenAndServe(url, nil))
}

func main() {
	token := os.Getenv("TOKEN")
	port := os.Getenv("PORT")
	go Listen(port)
	chatID, err := strconv.Atoi(os.Getenv("CHAT_ID"))
	if err != nil {
		log.Fatalln("Error env:", err)
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("Error bot creation", err)
	}
	log.Println("Access bot creation.")

	if _, err := bot.SetWebhook(tgbotapi.NewWebhook(webhook)); err != nil {
		log.Fatalf("Setting webhook %v", webhook)
	}

	updates := bot.ListenForWebhook("/")
	baseText := "Let Furgal free!\n"
	for update := range updates {
		text := fmt.Sprintf("%v%+v\n", baseText, update.Message)
		_, err := bot.Send(tgbotapi.NewMessage(int64(chatID), text))
		if err != nil {
			log.Printf("Error sending on %v", update)
		}
	}
}
