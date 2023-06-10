package factory

import "fmt"

type windowsDialog struct {
	button windowsButton
}

func NewWindowsDialog() windowsDialog {
	return windowsDialog{
		button: windowsButton{},
	}
}

func (d windowsDialog) Render() {
	fmt.Println("Windows Dialog Rendered")
	d.button.render()
}

func (d windowsDialog) ClickButton() {
	d.button.onClick()
}

// button
type windowsButton struct{}

func (b windowsButton) render() {
	fmt.Println("Windows Button Rendered")
}

func (b windowsButton) onClick() {
	fmt.Println("Click ! Windows Button Hello")
}
