package main

type operation struct {
	ID       int    `json:"id"`
	Position string `json:"position"`
	Cost     int    `json:"cost"`
	Date     int    `json:"date"`
}

type balance struct {
	Position string `json:"position"`
	Bill     int    `json:"bill"`
}
