package trigger

type BasicStringMatchTrigger struct {
	Match string
}

func (trigger BasicStringMatchTrigger) Check(str string) bool {
	return trigger.Match == str
}
