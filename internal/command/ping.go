package command

import (
	"bingoBotGo/internal/bot"
	"bingoBotGo/internal/trigger"
	"github.com/bwmarrin/discordgo"
	"log"
)

func MakePing() TriggeredCommand {
	return TriggeredCommand{
		Name:           "PingCommand",
		SelfTriggering: false,
		Trigger:        trigger.BasicStringMatch{Match: "ping"},
		Action:         pingAction,
	}
}

func pingAction(bot bot.IBot, session *discordgo.Session, message *discordgo.Message) (Result, error) {
	log.Println("Ping command triggered")

	_, err := session.ChannelMessageSend(message.ChannelID, "pong")
	if err != nil {
		log.Println("Failed to send message reply")
		return FAILURE, err
	}
	return SUCCESS, nil
}
