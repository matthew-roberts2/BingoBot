package command

import (
	"bingoBotGo/bot/command/trigger"
	"github.com/bwmarrin/discordgo"
	"log"
)

type PingCommand struct {
	Command
}

func MakePingCommand() PingCommand {
	return PingCommand{Command{trigger.BasicStringMatchTrigger{Match: "ping"}}}
}

func (command PingCommand) Process(bot IBot, session *discordgo.Session, message *discordgo.Message) Result {
	if !bot.IsSelf(message.Author) && command.Trigger.Check(message.Content) {
		_, err := session.ChannelMessageSend(message.ChannelID, "pong")
		if err != nil {
			log.Println("Failed to send message reply")
			return FAILURE
		}
		return SUCCESS
	}

	return PASS
}
