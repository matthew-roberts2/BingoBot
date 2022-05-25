package command

import (
	"bingoBotGo/bot/command/trigger"
	"github.com/bwmarrin/discordgo"
	"log"
)

type Ping struct {
	Command
}

func MakePing() Ping {
	return Ping{Command{trigger.BasicStringMatch{Match: "ping"}}}
}

func (command Ping) Process(bot IBot, session *discordgo.Session, message *discordgo.Message) Result {
	if !bot.IsSelf(message.Author) && command.Trigger.Check(message.Content, message.Author.ID) {
		log.Println("Ping command triggered")

		_, err := session.ChannelMessageSend(message.ChannelID, "pong")
		if err != nil {
			log.Println("Failed to send message reply")
			return FAILURE
		}
		return SUCCESS
	}

	return PASS
}
