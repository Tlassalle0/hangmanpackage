package hangmanpackage

import (
	"strconv"

	"github.com/rivo/tview"
)

// Updates all the boxes with the appropriate values
func UpdateBoxes(boxWord, boxLetter, boxOutput, boxHangman *tview.TextView, boxInput *tview.InputField, data HangManData, letterFile string) {
	if letterFile == "" {
		boxWord.SetText(data.Word + "\n" + strconv.Itoa(data.Attempts) + " Attempts remaining")
	} else {
		boxWord.SetText(SliceToStr(DisplayLetters(data.Word, letterFile)) + "\n" + strconv.Itoa(data.Attempts) + " Attempts remaining")
	}
	boxLetter.SetText(SliceToStr(data.UsedLetters))
	boxInput.SetText("")
	boxHangman.SetText(SliceToStr(DisplayJose(data.Attempts)))
}
