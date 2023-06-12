package product

import (
	"learndesignpattern/abstract_factory/shared/constant"
	"learndesignpattern/abstract_factory/shared/device"
)

type Macbook struct {
	Screen          constant.Screen
	OperationSystem constant.DesktopOS
}

func (n Macbook) ConnectToSmartphone(phone device.Smartphone) {
	phone.Connect(n)
}

func (n Macbook) ConnectToMonitor(monitor device.Monitor) {
	monitor.Connect(n)
}

func (n Macbook) GetOS() string {
	return string(n.OperationSystem)
}
