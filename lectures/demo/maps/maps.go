package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)

	shoppingList["Apple"] = 2
	shoppingList["Banana"] = 3
	shoppingList["Orange"] = 4

	shoppingList["Apple"] += 5

	fmt.Println(shoppingList)

	delete(shoppingList, "Banana")
	fmt.Println(shoppingList)
	fmt.Println(shoppingList["Banana"])

	banana, found := shoppingList["Banana"]
	fmt.Println(banana, found)

	apple, found := shoppingList["Apple"]
	fmt.Println(apple, found)

	var totalItems int
	for _, value := range shoppingList {
		totalItems += value
	}
	fmt.Println("Total items:", totalItems)
}
