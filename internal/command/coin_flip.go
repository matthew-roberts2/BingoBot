package command

import (
	"bingoBotGo/internal/bot"
	"bingoBotGo/internal/trigger"
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

func coinFlipAction(bot bot.IBot, session *discordgo.Session, message *discordgo.Message) (Result, error) {
	log.Println("Coin Flip command triggered")

	value := rand.Intn(2)

	response := "It's tails!"
	if value == 1 {
		response = "It's heads!"
	}

	_, err := session.ChannelMessageSend(message.ChannelID, response)
	if err != nil {
		log.Println("Failed to send message reply")
		return FAILURE, err
	}
	return SUCCESS, nil
}
