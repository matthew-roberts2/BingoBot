package trigger

type BasicStringMatch struct {
	Match string
}

func (trigger BasicStringMatch) Check(str string) bool {
	return trigger.Match == str
}
