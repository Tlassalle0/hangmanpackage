package hangman

import "strings"

func Game() {
	data, letterFile := GameStart(GetArgs())
	pointerData := &data
	gameOn := true
	app, flex, boxWord, boxLetter, boxOutput, boxInput, boxHangman := InitTview(data, letterFile)
	for gameOn {
		if err := app.SetRoot(flex, true).SetFocus(boxInput).Run(); err != nil {
			panic(err)
		}
		input := boxInput.GetText()
		if input == "STOP" {
			SaveFile(data)
		}
		if !CheckInput(input) {
			boxOutput.SetText("Invalid characters used")
			boxInput.SetText("")
			continue
		}
		input = strings.ToLower(input)
		var valid, alreadyUsed bool
		if len(input) == 1 {
			valid, alreadyUsed = AddLetter(pointerData, rune(input[0]))
		} else {
			valid = AddWord(pointerData, input)
		}
		if alreadyUsed {
			boxOutput.SetText("Letter already used")
		} else {
			if !valid {
				if len(input) == 1 {
					data.Attempts -= 1
				} else {
					data.Attempts -= 2
				}
				data.HangmanPositions = DisplayJose(data.Attempts)
				boxOutput.SetText("Oh no ! It's wrong !")
			} else {
				boxOutput.SetText("Great ! It's right !")
			}
		}
		UpdateBoxes(boxWord, boxLetter, boxOutput, boxHangman, boxInput, data, letterFile)
		gameOn = GameVerif(data, boxOutput)
	}
	endApp, endFlex, endButton := InitEnd(data)
	if err := endApp.SetRoot(endFlex, true).SetFocus(endButton).Run(); err != nil {
		panic(err)
	}
}
