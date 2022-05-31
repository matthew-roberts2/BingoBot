package mods

import (
	"bingoBotGo/internal/trigger"
	"github.com/bwmarrin/discordgo"
)

type UserGated struct {
	allowedUserIds map[string]struct{}
	trigger        trigger.Trigger
}

func MakeUserGated(trigger trigger.Trigger, allowedUserIds []string) UserGated {
	allowedUserIdsMap := make(map[string]struct{}, len(allowedUserIds))
	for _, usr := range allowedUserIds {
		allowedUserIdsMap[usr] = struct{}{}
	}

	return UserGated{allowedUserIdsMap, trigger}
}

func (trigger UserGated) Check(message *discordgo.Message) bool {
	_, allowed := trigger.allowedUserIds[message.Author.ID]
	if !allowed {
		return false
	}

	return trigger.trigger.Check(message)
}
