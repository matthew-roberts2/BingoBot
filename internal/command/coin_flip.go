package command

import (
	"bingoBotGo/internal/trigger"
	types "bingoBotGo/internal/types"
	"github.com/bwmarrin/discordgo"
	"log"
	"math/rand"
)

func MakeCoinFlip(botName string) TriggeredCommand {
	return TriggeredCommand{
		Name:           "CoinFlipCommand",
		SelfTriggering: false,
		Trigger:        trigger.MakeNamePrefixedBasicStringMatch(botName, "flip a coin"),
		Action:         coinFlipAction,
	}
}

func coinFlipAction(bot types.IBot, message *discordgo.Message) (Result, error) {
	log.Println("Coin Flip command triggered")

	value := rand.Intn(2)

	response := "It's tails!"
	if value == 1 {
		response = "It's heads!"
	}

	_, err := bot.SendMessageWithTyping(message.ChannelID, response)
	if err != nil {
		log.Println("Failed to send message reply")
		return FAILURE, err
	}
	return SUCCESS, nil
}
