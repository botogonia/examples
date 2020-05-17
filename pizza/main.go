package main

import (
	"log"
	"github.com/botogonia/flowbot"
)

func main() {
	bot, err := flowbot.NewFlowBot(
		"1256801395:AAENPIswnTiGpODZcucI6ycXcsfHumot8GY",
		60,
		"Вас долго не было, начните сначала")
	if err != nil {
		log.Panic(err)
	}
	//bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	bot.HandleUpdates(chatHandler3)
}
