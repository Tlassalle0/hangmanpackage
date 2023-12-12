package hangmanpackage

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func SaveFile(data HangManData) { // Save the current game state in a save.txt file
	print("File saved. Game exited")
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		print("Erreur de sauvegarde")
	}
	_ = ioutil.WriteFile("save.txt", file, 0644)
	os.Exit(0)
}
