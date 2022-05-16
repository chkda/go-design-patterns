package main

import "fmt"

type client struct {
}

func (c *client) insertLightningConnectorIntoComputer(com computer) {
	fmt.Println("Client insets Lightning Connector into Computer")
	com.insertIntoLightningPort()
}
