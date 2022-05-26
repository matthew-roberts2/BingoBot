package trigger

import "strings"

type PrefixVariantStringMatch struct {
	variants []string
}

func MakePrefixVariantStringMatch(variants []string) PrefixVariantStringMatch {
	return PrefixVariantStringMatch{variants}
}

func (trigger PrefixVariantStringMatch) Check(str string, _ string) bool {
	variantMatch := false

	for _, variant := range trigger.variants {
		variantMatch = variantMatch || strings.HasPrefix(str, variant)
	}

	return variantMatch
}
