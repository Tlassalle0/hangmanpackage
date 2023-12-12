package hangman

func CheckInput(input string) bool { //Checks if there isn't any non letter characters in a string
	for _, checkRune := range input {
		if checkRune < 65 || (checkRune > 90 && checkRune < 97) || checkRune > 122 {
			return false
		}
	}
	return true
}
