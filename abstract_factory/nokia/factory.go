package nokia

import (
	"learndesignpattern/abstract_factory/nokia/product"
	"learndesignpattern/abstract_factory/shared/constant"
	"learndesignpattern/abstract_factory/shared/device"
)

type Nokia struct{}

func (n Nokia) CreateMonitor() device.Monitor {
	return nil
}

func (n Nokia) CreateSmartphone() device.Smartphone {
	return product.NokiaPhone{
		Name: "Nokia Lumia X",
		Screen: constant.Screen{
			Size:        5,
			DisplayType: constant.Ips,
		},
		OperationSystem: constant.WindowsMobile,
	}
}

func (n Nokia) CreateNotebook() device.Notebook {
	return nil
}
