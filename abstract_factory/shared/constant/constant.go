package constant

import "fmt"

type Display string

const (
	Retina Display = "retina"
	Amoled Display = "amoled"
	Ips    Display = "IPS"
	Tn     Display = "TN"
)

type MobileOS string

const (
	Android       MobileOS = "android"
	IOS           MobileOS = "IOS"
	WindowsMobile MobileOS = "Windows Mobile"
)

type DesktopOS string

const (
	Windows DesktopOS = "windows"
	MacOS   DesktopOS = "MacOS"
	Linux   DesktopOS = "Linux"
)

type Screen struct {
	Size        uint8
	DisplayType Display
}

func (s Screen) Show(text string) {
	fmt.Println(text)
}

func (s Screen) ShowUnsupported(text string) {
	fmt.Println(text + " not supported")
}

func (s Screen) ConnectMessage(from string, to string) {
	fmt.Println(from + " connected to " + to)
}

const (
	Rotated     = "rotated"
	CantRotate  = "can't rotate device"
	Unsupported = "Device Not Supported"
)
