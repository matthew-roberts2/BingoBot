package command

import (
	"bingoBotGo/internal/bot/command/trigger"
	"github.com/bwmarrin/discordgo"
	"log"
	"math/rand"
)

type CoinFlip struct {
	Command
}

func MakeCoinFlip(botName string) CoinFlip {
	return CoinFlip{Command{trigger.MakeNamePrefixedBasicStringMatch(botName, "flip a coin")}}
}

func (command CoinFlip) Process(bot IBot, session *discordgo.Session, message *discordgo.Message) Result {
	if !bot.IsSelf(message.Author) && command.Trigger.Check(message.Content, message.Author.ID) {
		log.Println("Coin Flip command triggered")

		value := rand.Intn(2)

		response := "It's tails!"
		if value == 1 {
			response = "It's heads!"
		}

		_, err := session.ChannelMessageSend(message.ChannelID, response)
		if err != nil {
			log.Println("Failed to send message reply")
			return FAILURE
		}
		return SUCCESS
	}

	return PASS
}
