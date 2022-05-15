package main

import "fmt"

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")

	adidas_shoe := adidasFactory.makeShoe()
	adidas_shirt := adidasFactory.makeShirt()

	nike_shoe := nikeFactory.makeShoe()
	nike_shirt := nikeFactory.makeShirt()

	printShirtDetails(adidas_shirt)
	printShirtDetails(nike_shirt)

	printShoeDetails(adidas_shoe)
	printShoeDetails(nike_shoe)
}

func printShoeDetails(s iShoe) {
	fmt.Println("Shoe Logo:", s.getLogo())
	fmt.Println("Shoe Size:", s.getSize())
}

func printShirtDetails(s iShirt) {
	fmt.Println("Shirt Logo:", s.getLogo())
	fmt.Println("Shirt Size:", s.getSize())
}
