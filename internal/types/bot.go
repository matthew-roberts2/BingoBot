package bot

import "github.com/bwmarrin/discordgo"

type IBot interface {
	HandleMessage(session *discordgo.Session, message *discordgo.MessageCreate)
	IsSelf(userId string) bool
	GetInternalName() string
	GetGuildName(guildId string) string
	Session() *discordgo.Session
}
