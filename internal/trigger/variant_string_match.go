package trigger

import "github.com/bwmarrin/discordgo"

type VariantStringMatch struct {
	variants []string
}

func MakeVariantStringMatch(variants []string) VariantStringMatch {
	return VariantStringMatch{variants: variants}
}

func (trigger VariantStringMatch) Check(message *discordgo.Message) bool {
	variantMatch := false

	for _, variant := range trigger.variants {
		variantMatch = variantMatch || message.Content == variant
	}

	return variantMatch
}
