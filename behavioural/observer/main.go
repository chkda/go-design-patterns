package main

func main() {
	shirtItem := newItem("Nke Shirt")

	observer1 := &customer{id: "ads@gmail.com"}
	observer2 := &customer{id: "erp@gmail.com"}

	shirtItem.register(observer1)
	shirtItem.register(observer2)

	shirtItem.updateAvailability()
}
