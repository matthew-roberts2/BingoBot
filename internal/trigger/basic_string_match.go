package trigger

type BasicStringMatch struct {
	Match string
}

func (trigger BasicStringMatch) Check(str string, _ string) bool {
	return trigger.Match == str
}
