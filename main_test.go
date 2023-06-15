package main_test

import (
	"fmt"
	absfactory "learndesignpattern/abstract_factory"
	"learndesignpattern/builder"
	"learndesignpattern/factory"
	"learndesignpattern/prototype/document"
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

func TestBuilder(t *testing.T) {
	dir := builder.Director{}
	dir.SetBuilder(&builder.FamilyHouseBuilder{})
	house, err := dir.BuildMommmyHouse()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(house)
}

func TestPrototype(t *testing.T) {
	myPhoneBook := document.NewPhoneBook("Taufiq")
	copyPhoneBook := myPhoneBook.Clone()

	humairaPhoneBook, ok := copyPhoneBook.(*document.PhoneBook)
	if ok {
		humairaPhoneBook.BookOwner = "Humaira"
	}

	myPhoneBook.Print()
	humairaPhoneBook.Print()
}
