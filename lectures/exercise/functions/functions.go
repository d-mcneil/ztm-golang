package main

import "fmt"

//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greeting(name string) {
	fmt.Println("Hello,", name)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func whatsUp() string {
	return "What's up?"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func sum(a int, b int, c int) int {
	return a + b + c
}

//* Write a function that returns any number
func five() int {
	return 5
}

//* Write a function that returns any two numbers
func twoTwos() (int, int) {
	return 2, 2
}

//* Add three numbers together using any combination of the existing functions.
//  * Print the result
func addThree() {
	num1, num2 := twoTwos()
	num3 := five()
	fmt.Println(sum(num1, num2, num3))
}

func main() {
	//* Call every function at least once
	greeting("John")
	fmt.Println(whatsUp())
	addThree()
}
