package command

import (
	"bingoBotGo/internal/trigger"
	"bingoBotGo/internal/trigger/mods"
	types "bingoBotGo/internal/types"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

var matchVariants = []string{
	"I am ",
	"I'm ",
}

func MakeDadJoke() TriggeredCommand {
	return TriggeredCommand{
		Name:           "DadJokeCommand",
		SelfTriggering: false,
		Trigger:        mods.MakeRandomized(trigger.MakeVariantStringMatch(matchVariants), 0.15),
		Action:         dadJokeAction,
	}
}

func dadJokeAction(bot types.IBot, message *discordgo.Message) (Result, error) {
	log.Println("Dad Joke command triggered")

	var trimAmt int
	for _, variant := range matchVariants {
		if strings.HasPrefix(message.Content, variant) {
			trimAmt = len(variant)
			break
		}
	}
	injectStr := message.Content[trimAmt:]

	botName := bot.GetGuildName(message.GuildID)

	_, err := bot.Session().ChannelMessageSend(message.ChannelID, fmt.Sprintf("Hi %s, I'm %s", injectStr, botName))
	if err != nil {
		log.Println("Failed to send reply message")
		return FAILURE, err
	}

	return SUCCESS, nil
}
