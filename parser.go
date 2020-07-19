package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseMessage(text string) *Operation {
	var cost int
	DATA.ID++
	arr := strings.Split(text, " ")
	cost, err := strconv.Atoi(arr[1])
	if err != nil {
		cost = 0
	}
	date := time.Now().Unix()
	o := Operation{
		ID:       DATA.ID,
		Position: arr[0],
		Cost:     cost,
		Date:     int(date),
	}
	return &o
}

func Sum(o *Operation) {
	for ind, v := range DATA.Bill.List {
		if o.Position == v.Position {
			DATA.Bill.List[ind].FullCost += o.Cost
			return
		}
	}
	DATA.Bill.List = append(DATA.Bill.List, Balance{
		Position: o.Position,
		FullCost: o.Cost,
	})
}

func ParserPipeLine(text string) {
	o := ParseMessage(text)
	Sum(o)
}

func PrintBill() {
	bill := fmt.Sprintf("%+v", DATA.Bill)
	Send(bill)
}