package main

import (
	"bingoBotGo/internal/bot"
	"bingoBotGo/internal/util"
	"log"
)

func main() {
	util.LoadEnvs()

	bot.RunBingoBot()

	log.Println("App closed")
}
