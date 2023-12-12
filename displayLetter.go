package hangman

import (
	"bufio"
	"log"
	"os"
)

func DisplayLetters(word, letterFile string) []string { //Returns a slice of strings which represents the word in ascii art, using the "letterfile" file
	var lines []string
	for x := 0; x < 9; x++ {
		lines = append(lines, "")
	}
	for _, letter := range word {
		file, err := os.Open(letterFile)
		if err != nil {
			log.Fatal(1)
		}
		scanner := bufio.NewScanner(file)
		compt := 0
		line := 0
		for scanner.Scan() {
			if compt >= int(rune(letter-32))*9 && compt < int(rune(letter-31))*9 {
				lines[line] += scanner.Text()
				line += 1
			}
			compt += 1
		}
		file.Close()
	}
	for index := range lines {
		lines[index] += "\n"
	}
	return lines
}
