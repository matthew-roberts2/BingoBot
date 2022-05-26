package bot

import (
	"bingoBotGo/internal/bot/command"
	"github.com/bwmarrin/discordgo"
	"log"
)

type Bot struct {
	HumanName          string
	registeredCommands []command.ICommand
	Client             *discordgo.Session
}

type IBot interface {
	HandleMessage(session *discordgo.Session, message *discordgo.MessageCreate)
	IsSelf(author *discordgo.User) bool
	GetName() string
}

func MakeBot(client *discordgo.Session, name string) Bot {
	return Bot{registeredCommands: []command.ICommand{}, Client: client, HumanName: name}
}

func (bot *Bot) RegisterCommand(command command.ICommand) {
	log.Println("Registering command ", command)
	bot.registeredCommands = append(bot.registeredCommands, command)
}

func (bot *Bot) GetCommandCount() int {
	return len(bot.registeredCommands)
}

func (bot Bot) IsSelf(author *discordgo.User) bool {
	return author.ID == bot.Client.State.User.ID
}

func (bot Bot) GetName() string {
	return bot.HumanName
}

func (bot Bot) HandleMessage(session *discordgo.Session, messageCreate *discordgo.MessageCreate) {
	message := messageCreate.Message
	var commandResult command.Result
	for _, botCommand := range bot.registeredCommands {
		commandResult = botCommand.Process(bot, session, message)
		if commandResult == command.SUCCESS {
			break
		}
	}

	if commandResult != command.SUCCESS && !bot.IsSelf(messageCreate.Message.Author) {
		log.Println("Bot handler failed to find any matching commands")
	}
}
