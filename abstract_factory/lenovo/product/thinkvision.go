package product

import (
	"fmt"
	"learndesignpattern/abstract_factory/shared/constant"
	"learndesignpattern/abstract_factory/shared/device"
)

type ThinkVision struct {
	Name   string
	Screen constant.Screen
}

func (m ThinkVision) Connect(notebook device.Notebook) {
	switch notebook.GetOS() {
	case string(constant.Windows):
		m.Screen.ConnectMessage(m.Name, notebook.GetOS())
	case string(constant.Linux):
		m.Screen.ConnectMessage(m.Name, notebook.GetOS())
	case string(constant.MacOS):
		m.Screen.ConnectMessage(m.Name, notebook.GetOS())
	default:
		m.Screen.ShowUnsupported(notebook.GetOS())
	}
}

func (m ThinkVision) Rotate() {
	fmt.Println(constant.Rotated)
}
