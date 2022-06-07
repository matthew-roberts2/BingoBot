package trigger

import "github.com/bwmarrin/discordgo"

type Trigger interface {
	// Check - Checks whether the provided message triggers the Trigger
	Check(message *discordgo.Message) bool
}
