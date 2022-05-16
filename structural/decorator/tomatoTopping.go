package main

type tomatoTopping struct {
	pizza pizza
}

func (t *tomatoTopping) getPrice() int {
	price := t.pizza.getPrice()
	return price + 10
}
