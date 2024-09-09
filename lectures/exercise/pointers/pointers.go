//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type Product struct {
	name                string
	isSecurityTagActive bool
}

func activateSecurityTag(product *Product) {
	product.isSecurityTagActive = true
}

func deactivateSecurityTag(product *Product) {
	product.isSecurityTagActive = false
}

func checkout(productsPtr *[]*Product) {
	for _, productPtr := range *productsPtr {
		deactivateSecurityTag(productPtr)
	}
}

func printProducts(products *[]*Product) {
	fmt.Println("===== Products ===== ")
	for _, product := range *products {
		println(product.name, product.isSecurityTagActive)
	}
}

func main() {
	item1 := Product{"item1", true}
	item2 := Product{"item2", true}
	item3 := Product{"item3", true}
	item4 := Product{"item4", true}

	productsPtr := &[]*Product{&item1, &item2, &item3, &item4}
	printProducts(productsPtr)

	deactivateSecurityTag(&item2)
	deactivateSecurityTag((*productsPtr)[2])
	printProducts(productsPtr)

	checkout(productsPtr)
	printProducts(productsPtr)

	fmt.Println("===== Confirming that original items were modified, should all be false =====")
	fmt.Println(item1)
	fmt.Println(item2)
	fmt.Println(item3)
	fmt.Println(item4)
}
