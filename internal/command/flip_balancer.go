package command

import (
	"bingoBotGo/internal/trigger"
	types "bingoBotGo/internal/types"
	"github.com/bwmarrin/discordgo"
	"log"
)

type FlipBalancer struct {
	Command
}

const flipString = "(╯°□°）╯︵ ┻━┻"
const unflipString = "┬─┬ ノ( ゜-゜ノ)"

func MakeFlipBalancer() TriggeredCommand {
	return TriggeredCommand{
		Name:           "FlipBalancerCommand",
		SelfTriggering: false,
		Trigger:        trigger.MakeVariantStringMatch([]string{unflipString, flipString}),
		Action:         flipBalanceAction,
	}
}

func flipBalanceAction(bot types.IBot, message *discordgo.Message) (Result, error) {
	log.Println("Flip balance command triggered")

	responseString := flipString
	if message.Content == flipString {
		responseString = unflipString
	}

	_, err := bot.SendMessageWithTyping(message.ChannelID, responseString)
	if err != nil {
		log.Println("Failed to send message reply")
		return FAILURE, err
	}
	return SUCCESS, nil
}
