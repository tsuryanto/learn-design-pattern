package abstract_factory

import (
	"fmt"
	"learndesignpattern/abstract_factory/apple"
	"learndesignpattern/abstract_factory/lenovo"
	"learndesignpattern/abstract_factory/nokia"
	"learndesignpattern/abstract_factory/shared/device"
)

type GadgetFactory interface {
	CreateMonitor() device.Monitor
	CreateSmartphone() device.Smartphone
	CreateNotebook() device.Notebook
}

func GetGadgetFactory(brand string) (GadgetFactory, error) {
	if brand == "apple" {
		return &apple.Apple{}, nil
	}

	if brand == "lenovo" {
		return &lenovo.Lenovo{}, nil
	}

	if brand == "nokia" {
		return &nokia.Nokia{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}
