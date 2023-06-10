package factory

type Dialog interface {
	Render()
	ClickButton()
}

func GetDialog(osName string) Dialog {
	if osName == "windows" {
		return NewWindowsDialog()
	}
	return NewHtmlDialog()
}
