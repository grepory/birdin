package nms

import (
	"fmt"
	"math/rand"
)

var (
	ages = []string{
		"perpetual",
		"infant",
		"larva",
		"adult",
		"constant",
	}

	genders = []string{
		"exotic",
		"agender",
	}

	temperaments = []string{
		"sophisticated",
		"ambulatory",
		"resigned",
	}

	tweetFmt = `Age: %s
Gender: %s
Temperament: %s`
)

// Animal is an animal. Duh.
type Animal struct {
	Age         string
	Gender      string
	Temperament string
}

func (a Animal) String() string {
	return fmt.Sprintf(tweetFmt, a.Age, a.Gender, a.Temperament)
}

func genAnimal() Animal {
	return Animal{
		Age:         ages[rand.Int()%len(ages)],
		Gender:      genders[rand.Int()%len(genders)],
		Temperament: temperaments[rand.Int()%len(temperaments)],
	}
}

// Get a random animal tweet.
func Tweet() string {
	return genAnimal().String()
}
