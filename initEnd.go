package hangman

import "github.com/rivo/tview"

func InitEnd(data HangManData) (*tview.Application, *tview.Flex, *tview.Button) { // Initiates the end screen of the game
	app := tview.NewApplication()
	flex := tview.NewFlex()
	boxEnd := tview.NewTextView()
	if data.Attempts < 1 {
		boxEnd.SetText("░█░█░█▀█░█░█░░░█░░░█▀█░█▀▀░▀█▀\n░░█░░█░█░█░█░░░█░░░█░█░▀▀█░░█░\n░░▀░░▀▀▀░▀▀▀░░░▀▀▀░▀▀▀░▀▀▀░░▀░\nYour word was " + data.ToFind)
	} else {
		boxEnd.SetText("░█░█░█▀█░█░█░░░█░█░█▀█░█▀█\n░░█░░█░█░█░█░░░█▄█░█░█░█░█\n░░▀░░▀▀▀░▀▀▀░░░▀░▀░▀▀▀░▀░▀\nYour word was " + data.ToFind)
	}
	boxEnd.SetBorder(true)
	flex.AddItem(boxEnd, 0, 1, false)
	buttonEnd := tview.NewButton("Close").SetSelectedFunc(func() {
		app.Stop()
	})

	flex.AddItem(buttonEnd, 5, 1, true)
	return app, flex, buttonEnd
}
