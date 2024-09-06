//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var color = "black"
	year, age := 1900, 124
	// var (
	// 	first = 'X'
	// 	last = 'Y'
	// )
	var (
		first = "X"
		last = "Y"
	)
	var ageInDays int
	ageInDays = age * 365

	fmt.Println("My favorite color is", color)
	fmt.Println("I was born in", year, "and I am", age, "years old")
	fmt.Println("My initials are", first, last)
	fmt.Println("I am", ageInDays, "days old")
}