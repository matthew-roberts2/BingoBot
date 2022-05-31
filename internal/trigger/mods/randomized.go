package mods

import (
	"bingoBotGo/internal/trigger"
	"github.com/bwmarrin/discordgo"
	"log"
	"math/rand"
)

type Randomized struct {
	probability float64
	other       trigger.Trigger
}

// MakeRandomized - Constructs a Randomized version of the provided trigger. This trigger will first check if the
//                 contained trigger will fire, then checks a random number to see if it continues the trigger or
//                 fails it.
// probabilityPercent - The likelihood that the trigger will fire, as a floating point value from 0 to 1. Setting to 0 will still trigger, albeit extremely rarely
func MakeRandomized(other trigger.Trigger, probability float64) Randomized {
	return Randomized{probability, other}
}

func (trigger Randomized) Check(message *discordgo.Message) bool {
	randomSuccess := rand.Float64() <= trigger.probability
	otherSuccess := trigger.other.Check(message)

	if otherSuccess && !randomSuccess {
		log.Println("Trigger would have fired, but randomly decided not to")
	}

	return otherSuccess && randomSuccess
}
