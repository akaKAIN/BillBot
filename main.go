package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Listen() {
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	go Listen()
	token := os.Getenv("TOKEN")
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
	for update := range updates {
		_, err := bot.Send(tgbotapi.NewMessage(int64(chatID), "Let Furgal free!"))
		if err != nil {
			log.Printf("Error sending on %v", update)
		}
	}
}
