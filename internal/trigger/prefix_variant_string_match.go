package trigger

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type PrefixVariantStringMatch struct {
	variants []string
}

func MakePrefixVariantStringMatch(variants []string) PrefixVariantStringMatch {
	return PrefixVariantStringMatch{variants}
}

func (trigger PrefixVariantStringMatch) Check(message *discordgo.Message) bool {
	variantMatch := false

	for _, variant := range trigger.variants {
		variantMatch = variantMatch || strings.HasPrefix(message.Content, variant)
	}

	return variantMatch
}
