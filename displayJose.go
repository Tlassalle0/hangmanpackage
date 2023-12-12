package hangman

import (
	"bufio"
	"log"
	"os"
)

func DisplayJose(pos int) []string { //Returns a slice of strings which represents the state of the hanged man
	pos = 10 - pos
	file, err := os.Open("hangman.txt")
	if err != nil {
		log.Fatal(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	compt := 0

	for scanner.Scan() {
		if compt >= pos*8 && compt < (pos+1)*8 {
			lines = append(lines, scanner.Text()+"\n")
		}
		compt += 1
	}
	return lines
}
