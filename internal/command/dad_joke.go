package command

import (
	trigger2 "bingoBotGo/internal/trigger"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

var matchVariants = []string{
	"I am ",
	"I'm ",
}

type DadJoke struct {
	Command
}

func MakeDadJoke() DadJoke {
	return DadJoke{
		Command{
			trigger2.MakeRandomized(trigger2.MakePrefixVariantStringMatch(matchVariants), 0.1),
		},
	}
}

func (command DadJoke) Process(bot IBot, session *discordgo.Session, message *discordgo.Message) Result {
	if !bot.IsSelf(message.Author) && command.Trigger.Check(message.Content, message.Author.ID) {
		log.Println("Dad command triggered")

		trimAmt := -1
		for _, variant := range matchVariants {
			if strings.HasPrefix(message.Content, variant) {
				trimAmt = len(variant)
			}
		}
		injectStr := message.Content[trimAmt:]

		member, err := session.GuildMember(message.GuildID, session.State.User.ID)
		if err != nil {
			log.Println("Failed to look up self")
			return FAILURE
		}

		_, err = session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Hi %s, I'm %s", injectStr, member.Nick))
		if err != nil {
			log.Println("Failed to send message reply")
			return FAILURE
		}
		return SUCCESS
	}

	return PASS
}