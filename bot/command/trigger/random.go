package trigger

import (
	"log"
	"math/rand"
)

type Randomized struct {
	probabilityPercent float64
	other              Trigger
}

// MakeRandomized - Constructs a Randomized version of the provided trigger. This trigger will first check if the
//                 contained trigger will fire, then checks a random number to see if it continues the trigger or
//                 fails it.
// probabilityPercent - The likelihood that the trigger will fire from 0 to 1 as a percentage
func MakeRandomized(other Trigger, probabilityPercent float64) Randomized {
	return Randomized{probabilityPercent, other}
}

func (trigger Randomized) Check(str string, userId string) bool {
	randomSuccess := rand.Float64() <= trigger.probabilityPercent
	otherSuccess := trigger.other.Check(str, userId)

	if otherSuccess && !randomSuccess {
		log.Println("Trigger would have fired, but randomly decided not to")
	}

	return otherSuccess && randomSuccess
}
