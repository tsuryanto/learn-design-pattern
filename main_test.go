package main_test

import (
	absfactory "learndesignpattern/abstract_factory"
	"learndesignpattern/factory"
	"testing"
)

func TestFactory(t *testing.T) {
	// osName := "windows"
	osName := "linux"
	dialog := factory.GetDialog(osName)
	dialog.Render()
	dialog.ClickButton()
}

func TestAbstractFactory(t *testing.T) {
	apple, _ := absfactory.GetGadgetFactory("apple")
	lenovo, _ := absfactory.GetGadgetFactory("lenovo")
	nokia, _ := absfactory.GetGadgetFactory("nokia")

	myMonitor := apple.CreateMonitor()
	myPhone := nokia.CreateSmartphone()
	myNotebook := lenovo.CreateNotebook()

	myNotebook.ConnectToMonitor(myMonitor)
	myNotebook.ConnectToSmartphone(myPhone)
}
