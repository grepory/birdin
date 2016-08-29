package nms

import (
	"fmt"
	"math/rand"
)

var (
	// Thank you No Man's Sky Wikia
	// http://nomanssky.wikia.com/wiki/Behaviour_(Fauna)#Temperament
	ages = []string{
		"perpetual",
		"infant",
		"larval",
		"adult",
		"constant",
		"undying",
		"adolescent",
		"eternal",
		"fledgling",
	}

	genders = []string{
		"exotic",
		"agender",
		"gaseous",
		"liquid",
		"fluid",
		"chromatic",
		"indeterminate",
		"vectored",
		"radical",
		"asymmetric",
		"rational",
		"symetric",
		"alpha",
		"none",
		"female",
		"male",
		"non-uniform",
		"orthogonal",
	}

	temperaments = []string{
		"sophisticated",
		"ambulatory",
		"resigned",
		"rambunxious",
		"sleepy",
		"flatulent",
		"anxious",
		"grumpy",
		"locquacious",
		"amenable",
		"cautious",
		"defensive",
		"distinctive",
		"fearful",
		"migratory",
		"passive",
		"submissive",
		"unconcerned",
		"unpredictable",
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
