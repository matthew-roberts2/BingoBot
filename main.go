package main

import (
	"bingoBotGo/bot"
	"bingoBotGo/bot/command"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func loadEnvs() {
	log.Println("Loading environment data")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment files")
	}
}

func buildBingoBot(client *discordgo.Session) bot.IBot {
	log.Println("Constructing bot")

	bingoBot := bot.MakeBot(client, "Bingo")

	log.Println("Registering bot commands")

	bingoBot.RegisterCommand(command.MakePing())
	bingoBot.RegisterCommand(command.MakeCoinFlip(bingoBot.GetName()))
	bingoBot.RegisterCommand(command.MakeFlipBalancer())
	bingoBot.RegisterCommand(command.MakeDadJoke())

	log.Println("Registered", bingoBot.GetCommandCount(), "commands")

	client.AddHandler(bingoBot.HandleMessage)

	return bingoBot
}

func main() {
	loadEnvs()

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
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	log.Println("Received shutdown signal. Closing client session")
	err = client.Close()
	if err != nil {
		log.Println("Error closing client: ", err)
	}
}
