package adapter

// adapter for laptop
type VGAConnectorToHdmiConnectorAdapter struct {
	device OldDevice
}

func NewVGAConnectorToHdmiConnectorAdapter(device OldDevice) VGAConnectorToHdmiConnectorAdapter {
	return VGAConnectorToHdmiConnectorAdapter{
		device: device,
	}
}

func (v VGAConnectorToHdmiConnectorAdapter) ConnectByHdmiConnector() {
	v.device.ConnectByVGAConnector()
}
