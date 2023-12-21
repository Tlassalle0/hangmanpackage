package hangmanpackage

import (
	"math"
	"math/rand"
	"slices"
)

func RevealLetter(word string) []rune { // Reveals the right amount of letters in the word
	uniqueLetters := []rune{}
	for _, letter := range word {
		if !slices.Contains(uniqueLetters, letter) {
			uniqueLetters = append(uniqueLetters, letter)
		}
	}
	nbToReveal := int(math.Round(float64(len(uniqueLetters))/2.0 - 1.0))
	lettersToReveal := []rune{}
	for compt := 0; compt < nbToReveal; compt++ {
		randLetter := rand.Intn(len(uniqueLetters))
		lettersToReveal = append(lettersToReveal, uniqueLetters[randLetter])
		uniqueLetters = slices.Delete(uniqueLetters, randLetter, randLetter+1)
	}
	return lettersToReveal
}
