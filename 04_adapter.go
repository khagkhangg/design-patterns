package main

import "fmt"

type Client struct{}

type Mac1 struct{}

type Windows1 struct{}

type Computer1 interface {
	InsertIntoLightningPort()
}

type WindowsAdapter struct {
	windowMachine *Windows1
}

func (w *Windows1) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

func (m *Mac1) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer1) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}

func main() {

	client := &Client{}
	mac := &Mac1{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows1{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
	//Client inserts Lightning connector into computer.
	//	Lightning connector is plugged into mac machine.
	//	Client inserts Lightning connector into computer.
	//	Adapter converts Lightning signal to USB.
	//	USB connector is plugged into windows machine.

}
