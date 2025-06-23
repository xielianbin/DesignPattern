package main

import "adapter/internal"

func main() {
	client := &internal.Client{}
	mac := &internal.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &internal.Windows{}
	windowsMachineAdapter := &internal.WindowsAdapter{
		WindowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
