package main

import "fmt"

func main() {
	normal_builder, _ := getBuilder("normal")
	igloo_builder, _ := getBuilder("igloo")

	director := newDirector(normal_builder)
	normal_house := director.buildHouse()

	fmt.Println("Normal House Door:", normal_house.doorType)
	fmt.Println("Normal House Window:", normal_house.windowType)
	fmt.Println("Normal House Floors:", normal_house.floor)

	director.setBuilder(igloo_builder)
	igloo_house := director.buildHouse()

	fmt.Println("Igloo House Door:", igloo_house.doorType)
	fmt.Println("Igloo House Window:", igloo_house.windowType)
	fmt.Println("Igloo House Floors:", igloo_house.floor)

}
