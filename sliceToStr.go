package hangmanpackage

func SliceToStr(slice []string) string { // Converts a slice of string into a single string with spaces inbetween
	concatenated := ""
	for index, letter := range slice {
		concatenated += letter
		if index != len(slice) {
			concatenated += " "
		}
	}
	return concatenated
}
