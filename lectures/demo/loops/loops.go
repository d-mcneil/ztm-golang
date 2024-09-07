package main

import "fmt"

func main() {
	var sum int
	fmt.Println("Sum is", sum)
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println("Sum is", sum)
	}

	for sum > 10 {
		sum -= 5
		fmt.Println("Decrement. Sum is", sum)
	}

	for {
		sum -= 1
		fmt.Println("Decrement2. Sum is", sum)
		if sum == 0 {
			break
		}
	}
}
