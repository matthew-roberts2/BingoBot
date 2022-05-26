package main

import (
	"bingoBotGo/internal/bot"
	"bingoBotGo/internal/util"
	"log"
	"sync"
)

func main() {
	util.LoadEnvs()

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
