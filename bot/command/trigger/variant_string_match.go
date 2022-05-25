package trigger

type VariantStringMatch struct {
	variants []string
}

func MakeVariantStringMatch(variants []string) VariantStringMatch {
	return VariantStringMatch{variants: variants}
}

func (trigger VariantStringMatch) Check(str string, _ string) bool {
	variantMatch := false

	for _, variant := range trigger.variants {
		variantMatch = variantMatch || str == variant
	}

	return variantMatch
}
