package hangmanpackage

func AddLetter(data *HangManData, letter rune) (bool, bool) { //Checks if a letter is already used, and if it fits in the word to guess
	valid := false
	alreadyUsed := false
	for _, usedLetter := range data.UsedLetters {
		if string(letter) == usedLetter {
			alreadyUsed = true
		}
	}
	if !alreadyUsed {
		data.UsedLetters = append(data.UsedLetters, string(letter))
	}
	for index, letterToFind := range data.ToFind {
		if letterToFind == letter {
			newword := ""
			valid = true
			for indexInWord := range data.Word {
				if indexInWord == index {
					newword += string(letterToFind)
				} else {
					newword += string(data.Word[indexInWord])
				}
			}
			data.Word = newword
		}
	}
	return valid, alreadyUsed
}
