//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

func classifyAge(age int) {
	var classification string

	switch {
	case age == 0:
		classification = "newborn"
	case age <= 3:
		classification = "toddler"
	case age <= 12:
		classification = "child"
	case age <= 17:
		classification = "teenager"
	default:
		classification = "adult"
	}

	fmt.Println(classification)

}

func main() {
	classifyAge(6)
}
