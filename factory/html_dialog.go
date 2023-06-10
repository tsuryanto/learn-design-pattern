package factory

import "fmt"

type htmlDialog struct {
	button htmlButton
}

func NewHtmlDialog() htmlDialog {
	return htmlDialog{
		button: htmlButton{},
	}
}

func (d htmlDialog) Render() {
	fmt.Println("Html Dialog Rendered")
	d.button.render()
}

func (d htmlDialog) ClickButton() {
	d.button.onClick()
}

// button
type htmlButton struct{}

func (b htmlButton) render() {
	fmt.Println("Html Button Rendered")
}

func (b htmlButton) onClick() {
	fmt.Println("Click ! Html Button Hello")
}
