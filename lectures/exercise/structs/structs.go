//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	length, width int
}

func area(rectangle Rectangle) int {
	return rectangle.length * rectangle.width
}

func perimeter(rectangle Rectangle) int {
	return 2 * (rectangle.length + rectangle.width)
}

func main() {
	rectangle := Rectangle{length: 10, width: 5}

	fmt.Println("Area of rectangle:", area(rectangle))
	fmt.Println("Perimeter of rectangle:", perimeter(rectangle))

	rectangle.length *= 2
	rectangle.width *= 2

	fmt.Println("Area of rectangle:", area(rectangle))
	fmt.Println("Perimeter of rectangle:", perimeter(rectangle))
}
