package adapter

import "fmt"

type client struct {
}

func (c *client) insertLightningConnectorIntoComputer(com iComputer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.insertIntoLightningPort()
}
