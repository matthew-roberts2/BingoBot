package command

import (
	"bingoBotGo/internal/trigger"
	types "bingoBotGo/internal/types"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func MakeCountdown(botName string) TriggeredCommand {
	return TriggeredCommand{
		Name:           "CountdownCommand",
		SelfTriggering: false,
		Trigger:        trigger.MakeNamePrefixedBasicStringMatch(botName, "start a countdown"),
		Action:         countdownAction,
	}
}

func countdownAction(bot types.IBot, message *discordgo.Message) (Result, error) {
	go countdownRoutine(bot, message, 10)
	return SUCCESS, nil
}

func countdownRoutine(bot types.IBot, message *discordgo.Message, seconds int) {
	var err error
	var countdownMessage *discordgo.Message
	timeLeft := seconds

	// Defer a function to clean up the message once the countdown is finished
	defer func() {
		time.Sleep(5 * time.Second)
		if err := bot.Session().ChannelMessageDelete(countdownMessage.ChannelID, countdownMessage.ID); err != nil {
			log.Println("Failed to clean up countdown message")
		}
	}()

	for timeLeft >= 0 {
		var content string

		if timeLeft != 0 {
			content = fmt.Sprintf("%d", timeLeft)
		} else {
			content = ":stop_sign:"
		}

		if countdownMessage == nil {
			countdownMessage, err = bot.Session().ChannelMessageSend(message.ChannelID, content)
			if err != nil {
				log.Println("Failed to send countdown initialization message")
			}
		} else {
			countdownMessage, err = bot.Session().ChannelMessageEdit(countdownMessage.ChannelID, countdownMessage.ID, content)
			if err != nil {
				log.Println("Failed to send countdown message update")
			}
		}

		time.Sleep(1 * time.Second)
		timeLeft -= 1
	}
}
