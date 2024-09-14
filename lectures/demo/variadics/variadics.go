package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))                 // = 15
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) // = 55

	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 10}

	c := append(a, b...)
	fmt.Println(sum(c...)) // = 55
}
