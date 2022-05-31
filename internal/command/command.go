package command

import (
	"bingoBotGo/internal/bot"
	"bingoBotGo/internal/trigger"
	"github.com/bwmarrin/discordgo"
	"log"
)

type Command interface {
	Process(bot bot.IBot, session *discordgo.Session, message *discordgo.Message) Result
}

type Action func(bot bot.IBot, session *discordgo.Session, message *discordgo.Message) (Result, error)

type TriggeredCommand struct {
	Name           string
	SelfTriggering bool
	Trigger        trigger.Trigger
	Action         Action
}

func (command TriggeredCommand) Process(bot bot.IBot, session *discordgo.Session, message *discordgo.Message) Result {
	if (command.SelfTriggering || bot.IsSelf(message.Author)) && command.Trigger.Check(message) {
		if command.Action == nil {
			return FAILURE
		}

		result, err := command.Action(bot, session, message)
		if err != nil {
			log.Printf("Error processing %s command: %s", command.Name, err)
			return FAILURE
		}
		return result
	}
	return PASS
}

/*
                | Self Triggering | Not Self Triggering |
==========================================================
Msg is Self     | true            | false
Msg is not self | true            | true
*/

type Result int8

const (
	SUCCESS Result = iota
	FAILURE
	PASS
)
