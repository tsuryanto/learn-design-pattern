package product

import (
	"fmt"
	"learndesignpattern/abstract_factory/shared/constant"
	"learndesignpattern/abstract_factory/shared/device"
)

type AppleMonitor struct {
	Screen constant.Screen
}

func (m AppleMonitor) Connect(notebook device.Notebook) {
	switch notebook.GetOS() {
	case string(constant.MacOS):
		m.Screen.ConnectMessage("Apple Monitor", notebook.GetOS())
	default:
		m.Screen.ShowUnsupported(notebook.GetOS())
	}
}

func (m AppleMonitor) Rotate() {
	fmt.Println(constant.CantRotate)
}
