package main

import "fmt"

func double(number int) int {
	return number + number
}

func add(a, b int) int {
	return a + b
}

func greet() {
	fmt.Println("Hello, World!")
}

func main() {
	greet()
	
	dozen := double(6)
	fmt.Println("Double of 6 is", dozen)

	sum := add(3, 5)
	fmt.Println("3 + 5 =", sum)

	bakersDozen := dozen + 1
	fmt.Println("A baker's dozen is", bakersDozen)
	fmt.Println(add(double(6), 1))
}