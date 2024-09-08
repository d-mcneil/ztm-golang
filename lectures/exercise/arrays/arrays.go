//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Item struct {
	name  string
	price float64
}

func printLastItem(list [4]Item) {
	var lastItem Item
	for i := 0; i < len(list); i++ {
		item := list[i]
		if item.name != "" {
			lastItem = item
		}
	}
	if lastItem.name != "" {
		fmt.Println("Last item on the list:", lastItem.name)
	} else {
		fmt.Println("The list is empty")
	}
}

func printTotalPrice(list [4]Item) {
	var totalPrice float64
	for i := 0; i < len(list); i++ {
		item := list[i]
		totalPrice += item.price
	}
	fmt.Println("Total cost of the items:", totalPrice)
}

func printTotalItems(list [4]Item) {
	var totalItems int
	for i := 0; i < len(list); i++ {
		item := list[i]
		if item.name != "" {
			totalItems++
		}
	}
	fmt.Println("Total number of items:", totalItems)
}

func main() {

	shoppingList := [4]Item{}
	shoppingList[0] = Item{name: "Milk", price: 2.50}
	shoppingList[1] = Item{name: "Bread", price: 1.75}
	shoppingList[2] = Item{name: "Eggs", price: 3.99}

	printLastItem(shoppingList)
	printTotalPrice(shoppingList)
	printTotalItems(shoppingList)

	shoppingList[3] = Item{name: "Butter", price: 2.00}

	printLastItem(shoppingList)
	printTotalPrice(shoppingList)
	printTotalItems(shoppingList)
}
