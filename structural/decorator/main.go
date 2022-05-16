package main

import "fmt"

func main() {
	veg := &veggieMania{}

	vegWithCheeseTopping := &cheeseTopping{
		pizza: veg,
	}

	fmt.Println("Veggie Mania with Cheese Topping:", vegWithCheeseTopping.getPrice())

	vegWithCheeseAndTomatoTopping := &tomatoTopping{
		pizza: vegWithCheeseTopping,
	}

	fmt.Println("Veggie Mania with Cheese and Tomato Topping:", vegWithCheeseAndTomatoTopping.getPrice())
}
