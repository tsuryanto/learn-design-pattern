package apple

import (
	"learndesignpattern/abstract_factory/apple/product"
	"learndesignpattern/abstract_factory/shared/constant"
	"learndesignpattern/abstract_factory/shared/device"
)

type Apple struct{}

func (n Apple) CreateMonitor() device.Monitor {
	return product.AppleMonitor{
		Screen: constant.Screen{
			Size:        32,
			DisplayType: constant.Retina,
		},
	}
}

func (n Apple) CreateSmartphone() device.Smartphone {
	return product.IPhone{
		Screen: constant.Screen{
			Size:        6,
			DisplayType: constant.Amoled,
		},
		OperationSystem: constant.IOS,
	}
}

func (n Apple) CreateNotebook() device.Notebook {
	return product.Macbook{
		Screen: constant.Screen{
			Size:        13,
			DisplayType: constant.Retina,
		},
		OperationSystem: constant.MacOS,
	}
}
