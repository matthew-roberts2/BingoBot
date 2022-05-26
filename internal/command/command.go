package command

import (
	"bingoBotGo/internal/trigger"
	"github.com/bwmarrin/discordgo"
)

// IBot - Stand-in interface to avoid a circular dependency
type IBot interface {
	IsSelf(author *discordgo.User) bool
}

type Command struct {
	Trigger trigger.Trigger
}

type ICommand interface {
	Process(bot IBot, session *discordgo.Session, message *discordgo.Message) Result
}

type Result int8

const (
	SUCCESS Result = iota
	FAILURE
	PASS
)
