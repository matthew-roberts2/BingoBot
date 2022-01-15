package command

import (
	"bingoBotGo/bot/command/trigger"
	"github.com/bwmarrin/discordgo"
	"log"
)

type FlipBalancer struct {
	Command
}

const flipString = "(╯°□°）╯︵ ┻━┻"
const unflipString = "┬─┬ ノ( ゜-゜ノ)"

func MakeFlipBalancer() FlipBalancer {
	return FlipBalancer{Command{Trigger: trigger.MakeVariantStringMatch([]string{unflipString, flipString})}}
}

func (command FlipBalancer) Process(bot IBot, session *discordgo.Session, message *discordgo.Message) Result {
	if !bot.IsSelf(message.Author) && command.Trigger.Check(message.Content) {
		log.Println("Flip Balancer command triggered")

		responseString := flipString

		if message.Content == flipString {
			responseString = unflipString
		}

		_, err := session.ChannelMessageSend(message.ChannelID, responseString)
		if err != nil {
			log.Println("Failed to send message reply")
			return FAILURE
		}
		return SUCCESS
	}

	return PASS
}
