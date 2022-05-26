package main

import (
	"bingoBotGo/internal/bot"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

var _ = loadEnvs()

func loadEnvs() bool {
	log.Println("Loading environment data")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment files")
	}

	return true
}

func main() {
	var waitGroup sync.WaitGroup
	var funcs = []func(wg *sync.WaitGroup){
		bot.RunBingoBot,
	}

	for _, f := range funcs {
		waitGroup.Add(1)
		go f(&waitGroup)
	}

	log.Println("Prepared threads")
	waitGroup.Wait()
	log.Println("App closed")
}
