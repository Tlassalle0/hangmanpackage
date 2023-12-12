package hangmanpackage

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Initiates all the boxes for the game
func InitTview(data HangManData, letterFile string) (*tview.Application, *tview.Flex, *tview.TextView, *tview.TextView, *tview.TextView, *tview.InputField, *tview.TextView) {
	app := tview.NewApplication()

	// Creates all the boxes
	box1 := tview.NewTextView()
	box2 := tview.NewTextView()
	box3 := tview.NewTextView()
	box4 := tview.NewInputField().
		SetLabel("What's your guess ? : ").
		SetAcceptanceFunc(nil)

	box5 := tview.NewTextView()

	box4.SetDoneFunc(func(key tcell.Key) {
		app.Stop()
	})

	box1.SetTitle("Word and Attempts")
	box2.SetTitle("Used Letters")
	box3.SetTitle("Output")
	box4.SetTitle("Input")
	box5.SetTitle("Hangman")

	// Set text for the text views
	if letterFile == "" {
		box1.SetText(data.Word + "\n" + strconv.Itoa(data.Attempts) + " Attempts remaining")
	} else {
		box1.SetText(SliceToStr(DisplayLetters(data.Word, letterFile)) + "\n" + strconv.Itoa(data.Attempts) + " Attempts remaining")
	}
	box2.SetText(SliceToStr(data.UsedLetters))
	box3.SetText("The program will answer here")
	box5.SetText(SliceToStr(data.HangmanPositions))

	box1.SetBorder(true)
	box2.SetBorder(true)
	box3.SetBorder(true)
	box4.SetBorder(true)
	box5.SetBorder(true)

	// Set the text views' text alignment
	box1.SetTextAlign(tview.AlignCenter)
	box2.SetTextAlign(tview.AlignCenter)
	box3.SetTextAlign(tview.AlignCenter)
	box5.SetTextAlign(tview.AlignCenter)

	// Create a flex layout to arrange the components
	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(box1, 0, 1, false).
			AddItem(box2, 0, 1, false).
			AddItem(box3, 5, 1, false).
			AddItem(box4, 5, 1, false), 0, 1, false).
		AddItem(box5, 20, 0, false)
	return app, flex, box1, box2, box3, box4, box5
}
