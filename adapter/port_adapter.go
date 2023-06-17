package adapter

// adapter for External Display
type VGAPortToHdmiPortAdapter struct {
	projector ExternalDisplay
}

func NewVGAPortToHdmiPortAdapter(projector ExternalDisplay) VGAPortToHdmiPortAdapter {
	return VGAPortToHdmiPortAdapter{
		projector: projector,
	}
}

func (p VGAPortToHdmiPortAdapter) OpenHdmiPort(device NewDevice) {
	device.ConnectByHdmiConnector()
	p.projector.Display()
}
