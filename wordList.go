package hangmanpackage

import (
	"bufio"
	"log"
	"os"
)

type HangManData struct {
	Word             string   // Word composed of '_', ex: H_ll_
	ToFind           string   // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int      // Number of attempts left
	UsedLetters      []string //All the used letters
	HangmanPositions []string // It can be the array where the positions parsed in "hangman.txt" are stored
}

func GetWords(wordList string) []string { //Return a string of strings of all the lines of a file wordList
	file, err := os.Open(wordList)
	if err != nil {
		log.Fatal(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
