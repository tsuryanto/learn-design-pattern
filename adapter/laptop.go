package adapter

import "fmt"

type ExternalDisplay interface {
	Display()
}

type AcerSwift struct{}

func (a AcerSwift) ConnectByHdmiConnector() {
	fmt.Println("Request Connection to external display from Acer")
}

type Thinkbook struct{}

func (t Thinkbook) ConnectByVGAConnector() {
	fmt.Println("Request Connection to external display from Thinkbook")
}
