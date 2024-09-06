package main

import "fmt"

func main() {
	var myName = "Dean"	
	fmt.Println("My name is", myName) // Space is added automatically

	var name string = "Kathy"
	fmt.Println("name =", name)

	username := "admin"
	fmt.Println("username =", username)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 1, 5
	fmt.Println("part 1 is", part1, "other is", other)

	part2, other := 2, 10
	fmt.Println("part 2 is", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("sum = is", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)

	fmt.Println("lessonName=", lessonName)
	fmt.Println("lessonType=", lessonType)

	word1, word2, _ := "hello", "world", "!" // exclamation mark is ignored because of the underscore
	fmt.Println(word1, word2,)
}