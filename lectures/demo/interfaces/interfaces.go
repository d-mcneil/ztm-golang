package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("Cooking chicken")
}

func (s Salad) PrepareDish() {
	fmt.Println("Chopping salad")
	fmt.Println("Adding dressing")
}

func prepareDishes(dishes []Preparer) {
	for _, dish := range dishes {
		fmt.Printf("--Dish: %v--\n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}

func main() {
	dishes := []Preparer{
		Chicken("Chicken"),
		Salad("Salad"),
	}

	prepareDishes(dishes)
}
