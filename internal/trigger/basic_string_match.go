package trigger

import "github.com/bwmarrin/discordgo"

type BasicStringMatch struct {
	Match string
}

func (trigger BasicStringMatch) Check(message *discordgo.Message) bool {
	return trigger.Match == message.Content
}
