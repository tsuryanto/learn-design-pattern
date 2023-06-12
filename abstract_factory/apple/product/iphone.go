package product

import (
	"learndesignpattern/abstract_factory/shared/constant"
	"learndesignpattern/abstract_factory/shared/device"
)

type IPhone struct {
	Screen          constant.Screen
	OperationSystem constant.MobileOS
}

func (s IPhone) Connect(notebook device.Notebook) {
	switch notebook.GetOS() {
	case string(constant.MacOS):
		s.Screen.ConnectMessage("Iphone", notebook.GetOS())
	default:
		s.Screen.ShowUnsupported(notebook.GetOS())
	}
}

func (s IPhone) OpenCamera() {
	s.Screen.Show(string(s.OperationSystem) + " Opening the Camera")
}
