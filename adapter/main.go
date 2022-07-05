// The objective of the Adapter is to create acts as a middle method to allow adaptability between two functionalities.

package main

import "fmt"

type client struct {
}

func (c *client) insertLightningConnectorIntoComputer(com computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.insertIntoLightningPort()
}

type computer interface {
	insertIntoLightningPort()
}

type mac struct {
}

func (m *mac) insertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type windows struct{}

func (w *windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

type windowsAdapter struct {
	windowsMachine *windows
}

func (w *windowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowsMachine.insertIntoUSBPort()
}

func main() {
	client := &client{}

	// Lightning in MAC
	mac := &mac{}
	client.insertLightningConnectorIntoComputer(mac)

	// Lightning in Windows using Adapter
	windows := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowsMachine: windows,
	}

	client.insertLightningConnectorIntoComputer(windowsMachineAdapter)
}
