package main_test

import (
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
