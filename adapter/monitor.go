package adapter

import "fmt"

type NewDevice interface {
	ConnectByHdmiConnector()
}

type Monitor struct {
}

func (m Monitor) OpenHdmiPort(device NewDevice) {
	device.ConnectByHdmiConnector()
	m.Display()
}

func (m Monitor) Display() {
	fmt.Println("Monitor connected")
}
