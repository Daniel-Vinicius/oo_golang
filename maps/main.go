package main

import "fmt"

func main() {
	var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	alphabetByNumbers := make(map[string]int)

	alphabetIndex := len(alphabet) - 1

	for index := 0; index <= alphabetIndex; index++ {
		letter := alphabet[index]
		alphabetByNumbers[letter] = index + 1
	}

	// delete(alphabetByNumbers, "y")

	fmt.Println(alphabetByNumbers)
	// fmt.Println(len(alphabet))

	positionOfY, yExists := alphabetByNumbers["y"]
	if yExists {
		fmt.Println(positionOfY, "Y Exists!", yExists)
	}

	// var m = map[int]string{1: "a"}
	// fmt.Println(m)
}
