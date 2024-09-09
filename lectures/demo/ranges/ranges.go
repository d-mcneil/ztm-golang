package main

import "fmt"

func main() {
	word := "Ωa"
	for i := 0; i < len(word); i++ {
		fmt.Println(word[i])
		// THIS PRINTS 3 BYTES; 206, 169, AND 97; NOT CHARACTERS
		// range lets you access characters in a string
	}
	fmt.Println()
	for _, letter := range word {
		fmt.Println(letter)
		// THIS PRINTS TWO CHARACTERS, 937 AND 97
	}
	fmt.Println()
	// --------------------------------------------------------------------------------
	slice := []string{"Hello", "world", "!"}

	for i, element := range slice {
		fmt.Println(i, element, ":")
		for _, ch := range element {
			fmt.Printf("     %q\n", ch) // %q prints out glyph representation of the runes
		}
	}

}
