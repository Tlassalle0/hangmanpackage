package hangman

import (
	"math/rand"
	"strings"
)

func CreateData(words []string) HangManData {
	index := rand.Intn(len(words))
	goal := strings.ToLower(words[index])
	word := ""
	ltToReveal := RevealLetter(goal)
	for _, letter := range goal {
		inWord := false
		for _, checkReveal := range ltToReveal {
			if letter == checkReveal {
				word += string(letter)
				inWord = true
			}
		}
		if !inWord {
			word += "_"
		}
	}
	letters := make([]string, 0)
	for _, revealedLetter := range ltToReveal {
		letters = append(letters, string(revealedLetter))
	}
	data := HangManData{word, goal, 10, letters, DisplayJose(10)}
	return data
}
