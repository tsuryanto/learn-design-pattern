package product

import (
	"learndesignpattern/abstract_factory/shared/constant"
	"learndesignpattern/abstract_factory/shared/device"
)

type Thinkbook struct {
	Screen          constant.Screen
	OperationSystem constant.DesktopOS
}

func (n Thinkbook) ConnectToSmartphone(phone device.Smartphone) {
	phone.Connect(n)
}

func (n Thinkbook) ConnectToMonitor(monitor device.Monitor) {
	monitor.Connect(n)
}

func (n Thinkbook) GetOS() string {
	return string(n.OperationSystem)
}
