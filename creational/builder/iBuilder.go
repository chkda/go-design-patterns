package main

import "fmt"

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(builder string) (iBuilder, error) {
	switch builder {
	case "normal":
		return &normalBuilder{}, nil
	case "igloo":
		return &iglooBuilder{}, nil
	default:
		return nil, fmt.Errorf("Wrong choice")
	}
}
