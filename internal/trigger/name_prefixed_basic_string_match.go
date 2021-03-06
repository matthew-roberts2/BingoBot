package trigger

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var nameVariantFormats = [...]string{
	"Hey %s",
	"%s",
	"Yo %s",
}

const nameVariantCount = len(nameVariantFormats)

type NamePrefixedBasicStringMatch struct {
	rawName      string
	nameVariants [nameVariantCount]string
	textMatch    string
}

func MakeNamePrefixedBasicStringMatch(namePrefix string, textMatch string) NamePrefixedBasicStringMatch {
	var variants [nameVariantCount]string

	for i, variant := range nameVariantFormats {
		variants[i] = fmt.Sprintf(variant, namePrefix)
	}

	return NamePrefixedBasicStringMatch{
		rawName:      namePrefix,
		nameVariants: variants,
		textMatch:    textMatch,
	}
}

func (trigger NamePrefixedBasicStringMatch) Check(message *discordgo.Message) bool {
	str := message.Content
	nameVariantMatch := false

	for _, variant := range trigger.nameVariants {
		nameVariantMatch = nameVariantMatch || strings.HasPrefix(str, variant)
	}

	return nameVariantMatch && strings.Contains(str, trigger.textMatch)
}
