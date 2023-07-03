package main_test

import (
	"fmt"
	absfactory "learndesignpattern/abstract_factory"
	"learndesignpattern/adapter"
	"learndesignpattern/bridge"
	"learndesignpattern/builder"
	"learndesignpattern/factory"
	"learndesignpattern/prototype/document"
	"learndesignpattern/singleton"
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

func TestSingleton(t *testing.T) {
	dbWithLock := singleton.GetDBInsstanceWithLock()
	fmt.Println(*dbWithLock) // set instance, reeturn instance
	fmt.Println(*dbWithLock) // return instance only

	elastic := singleton.GetElasticInstanceWithOnce()
	fmt.Println(*elastic) // set instance, reeturn instance
	fmt.Println(*elastic) // return instance only
}

func TestAdapter(t *testing.T) {
	// laoptop with HDMI Port only
	myLaptop := adapter.AcerSwift{}
	myOldLaptop := adapter.Thinkbook{}

	monitor := adapter.Monitor{}
	projector := adapter.Infocus{}

	// my laptop connect to monitor
	monitor.OpenHdmiPort(myLaptop)

	// my laptop connect to projector
	projectorHdmiAdapter := adapter.NewVGAPortToHdmiPortAdapter(projector)
	projectorHdmiAdapter.OpenHdmiPort(myLaptop)

	// my old laptop connecto to projector
	projector.OpenVgaPort(myOldLaptop)

	// my old laptop to monitor
	myOldLaptopHdmiAdapter := adapter.NewVGAConnectorToHdmiConnectorAdapter(myOldLaptop)
	monitor.OpenHdmiPort(myOldLaptopHdmiAdapter)
}

func TestBridge(t *testing.T) {
	clockGreen := bridge.Cycle{
		Radius: 7,
		Color:  bridge.Green{},
	}

	tissueRed := bridge.Rectangle{
		SideLength: 10,
		Color:      bridge.Red{},
	}

	fmt.Println("Area :", clockGreen.Area())
	fmt.Println("Area :", tissueRed.Area())

	clockGreen.Draw()
	tissueRed.Draw()

	clockRed := clockGreen
	clockRed.Color = bridge.Red{}
	clockRed.Draw()

	tissueGreen := tissueRed
	tissueGreen.Color = bridge.Green{}
	tissueGreen.Draw()
}
