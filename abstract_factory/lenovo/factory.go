package lenovo

import (
	"learndesignpattern/abstract_factory/lenovo/product"
	"learndesignpattern/abstract_factory/shared/constant"
	"learndesignpattern/abstract_factory/shared/device"
)

type Lenovo struct{}

func (n Lenovo) CreateMonitor() device.Monitor {
	return product.ThinkVision{
		Screen: constant.Screen{
			Size:        24,
			DisplayType: constant.Ips,
		},
	}
}

func (n Lenovo) CreateSmartphone() device.Smartphone {
	return product.LenovoPhone{
		Name: "Lenovo Android",
		Screen: constant.Screen{
			Size:        6,
			DisplayType: constant.Amoled,
		},
		OperationSystem: constant.Android,
	}
}

func (n Lenovo) CreateNotebook() device.Notebook {
	return product.Thinkbook{
		Screen: constant.Screen{
			Size:        14,
			DisplayType: constant.Tn,
		},
		OperationSystem: constant.Windows,
	}
}
