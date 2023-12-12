package hangman

import (
	"github.com/rivo/tview"
)

func GameVerif(data HangManData, boxOutput *tview.TextView) bool { //Checks if the game is not over
	if data.Word == data.ToFind {
		return false
	}
	if data.Attempts < 1 {
		return false
	}
	return true
}
