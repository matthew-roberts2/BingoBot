package trigger

type Trigger interface {
	// Check - Checks whether the provided message triggers the Trigger
	Check(str string) bool
}
