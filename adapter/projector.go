package adapter

import "fmt"

type OldDevice interface {
	ConnectByVGAConnector()
}

type Infocus struct {
}

func (i Infocus) OpenVgaPort(device OldDevice) {
	device.ConnectByVGAConnector()
	i.Display()
}

func (i Infocus) Display() {
	fmt.Println("Infocus connected")
}
