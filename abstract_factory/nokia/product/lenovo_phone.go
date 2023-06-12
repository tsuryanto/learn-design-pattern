package product

import (
	"learndesignpattern/abstract_factory/shared/constant"
	"learndesignpattern/abstract_factory/shared/device"
)

type NokiaPhone struct {
	Name            string
	Screen          constant.Screen
	OperationSystem constant.MobileOS
}

func (s NokiaPhone) Connect(notebook device.Notebook) {
	switch notebook.GetOS() {
	case string(constant.Windows):
		s.Screen.ConnectMessage(s.Name, notebook.GetOS())
	case string(constant.Linux):
		s.Screen.ConnectMessage(s.Name, notebook.GetOS())
	case string(constant.MacOS):
		s.Screen.ConnectMessage(s.Name, notebook.GetOS())
	default:
		s.Screen.ShowUnsupported(notebook.GetOS())
	}
}

func (s NokiaPhone) OpenCamera() {
	s.Screen.Show(string(s.OperationSystem) + " Opening the Camera")
}
