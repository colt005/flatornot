package common

import (
	"math/big"

	"crypto/rand"
)

func CalculatePercentage(part, total int) float64 {
	if total == 0 {
		return 0
	}
	return float64(part) / float64(total) * 100
}

func Mul(a, b int) int {
	return a * b
}

type Puns map[string][]string

var P = Puns{
	"Flat": {
		`Whoa, watch out for the edge! Don’t fall off!`,
		`Flat Earthers unite! NASA is quaking right now!`,
		`Please don't tell Columbus, he had a tough time.`,
		`Alert: We have a time traveler from the 15th century!`,
		`Congratulations! You’ve just unlocked membership to the Flat Earth Society!`,
	},
	"Round": {
		`Welcome to the 21st century, where physics and science still rule!`,
		`Thank you for not sailing off the edge of reason.`,
		`Nice! Now let's try convincing people about climate change next!`,
		`Phew, we were starting to lose faith in gravity!`,
		`Good job! You passed the easiest science test ever!`,
	},
	"Waiting": {
		`Come on, it’s either a pancake or a beach ball! Pick a side!`,
		`Still deciding? Just look out of your window… Or a plane, maybe?`,
		`It’s ok to be confused, the universe is a mysterious place!`,
		`Can’t choose? That’s ok, you’re in the Bermuda Triangle of decisions.`,
		`Undecided? Maybe try spinning a globe for some clarity!`,
	},
}

func (p *Puns) Random(t string) string {
	puns, ok := (*p)[t]
	if !ok {
		return "Invalid category!"
	}
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(puns))))
	if err != nil {
		return "Error generating random number!"
	}
	randomIndex := int(nBig.Int64())
	return puns[randomIndex]

}
