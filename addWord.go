package hangmanpackage

func AddWord(data *HangManData, word string) bool { //Checks if the word is equal to the goal of the game
	found := false
	if word == data.ToFind {
		data.Word = word
		found = true
	}
	return found
}
