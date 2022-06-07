package command

import (
	"bingoBotGo/internal/trigger"
	types "bingoBotGo/internal/types"
	"github.com/bwmarrin/discordgo"
	"log"
)

type Command interface {
	GetName() string
	Process(bot types.IBot, message *discordgo.Message) Result
}

type Action func(bot types.IBot, message *discordgo.Message) (Result, error)

type TriggeredCommand struct {
	Name           string
	SelfTriggering bool
	Trigger        trigger.Trigger
	Action         Action
}

func (command TriggeredCommand) GetName() string {
	return command.Name
}

func (command TriggeredCommand) Process(bot types.IBot, message *discordgo.Message) Result {
	if (command.SelfTriggering || bot.IsSelf(message.Author.ID)) && command.Trigger.Check(message) {
		if command.Action == nil {
			return FAILURE
		}

		result, err := command.Action(bot, message)
		if err != nil {
			log.Printf("Error processing %s command: %s", command.Name, err)
			return FAILURE
		}
		return result
	}
	return PASS
}

type Result int8

const (
	SUCCESS Result = iota
	FAILURE
	PASS
)
