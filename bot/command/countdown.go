package command

import (
	"bingoBotGo/bot/command/trigger"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

type Countdown struct {
	Command
}

func MakeCountdown(botName string) Countdown {
	return Countdown{Command{trigger.MakeNamePrefixedBasicStringMatch(botName, "start a countdown")}}
}

func countdown(session *discordgo.Session, message *discordgo.Message, seconds int) {
	var boundMessage *discordgo.Message
	haveBoundMessage := false
	timeLeft := seconds
	defer func() {
		time.Sleep(5 * time.Second)
		if err := session.ChannelMessageDelete(boundMessage.ChannelID, boundMessage.ID); err != nil {
			log.Println("Failed to clean up message")
		}
	}()
	for timeLeft >= 0 {
		var content string
		if timeLeft != 0 {
			content = fmt.Sprintf("%d", timeLeft)
		} else {
			content = ":stop_sign:"
		}

		if !haveBoundMessage {
			m, err := session.ChannelMessageSend(message.ChannelID, content)
			if err != nil {
				log.Println("Failed to send message")
			}
			boundMessage = m
			haveBoundMessage = true
		} else {
			m, err := session.ChannelMessageEdit(boundMessage.ChannelID, boundMessage.ID, content)
			if err != nil {
				log.Println("Failed to update message")
			}
			boundMessage = m
		}
		time.Sleep(1 * time.Second)
		timeLeft -= 1
	}
}

func (command Countdown) Process(bot IBot, session *discordgo.Session, message *discordgo.Message) Result {
	if !bot.IsSelf(message.Author) && command.Trigger.Check(message.Content, message.Author.ID) {
		go countdown(session, message, 10)
		return SUCCESS
	}
	return PASS
}
