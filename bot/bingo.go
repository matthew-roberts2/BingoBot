package bot

import (
	"bingoBotGo/bot/command"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func buildBingoBot(client *discordgo.Session) IBot {
	log.Println("Constructing bot")

	bingoBot := MakeBot(client, "Bingo")

	log.Println("Registering bot commands")

	bingoBot.RegisterCommand(command.MakePing())
	bingoBot.RegisterCommand(command.MakeCoinFlip(bingoBot.GetName()))
	bingoBot.RegisterCommand(command.MakeFlipBalancer())
	bingoBot.RegisterCommand(command.MakeDadJoke())
	bingoBot.RegisterCommand(command.MakeCountdown(bingoBot.GetName()))

	log.Println("Registered", bingoBot.GetCommandCount(), "commands")

	client.AddHandler(bingoBot.HandleMessage)

	return bingoBot
}

func RunBingoBot(wg *sync.WaitGroup) {
	botToken := os.Getenv("BINGO_BOT_TOKEN")

	if botToken == "" {
		log.Fatal("No bot token specified in environment")
	}

	log.Println("Building Discord client")

	client, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal("Error creating Discord client: ", err)
	}

	client.Identify.Intents = discordgo.IntentsGuildMessages

	_ = buildBingoBot(client)

	log.Println("Opening client session")
	err = client.Open()
	if err != nil {
		log.Fatal("Failed to open connection for client: ", err)
	}

	log.Println("Bot started")
	defer func(client *discordgo.Session) {
		log.Println("Closing Discord session")
		err := client.Close()
		if err != nil {
			log.Println("Failed to close discord client")
		}
	}(client)

	notify := make(chan os.Signal)
	signal.Notify(notify, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-notify

	log.Println("Done runBot")
	wg.Done()
}
