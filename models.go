package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type Operation struct {
	ID       int    `json:"id"`
	Position string `json:"position"`
	Cost     int    `json:"cost"`
	Date     int    `json:"date"`
}

type Balance struct {
	Position string `json:"position"`
	FullCost int    `json:"full_cost"`
}

type Bill struct {
	List []Balance
}

type Data struct {
	ID      int    `json:"id"`
	Bill    Bill   `json:"Bill"`
	TOKEN   string `json:"token"`
	PORT    string `json:"port"`
	CHAT_ID int    `json:"chat_id"`
	Bot     *tgbotapi.BotAPI
}
