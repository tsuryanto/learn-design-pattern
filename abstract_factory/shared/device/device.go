package device

type Monitor interface {
	Connect(notebook Notebook)
	Rotate()
}

type Smartphone interface {
	Connect(notebook Notebook)
	OpenCamera()
}

type Notebook interface {
	ConnectToSmartphone(phone Smartphone)
	ConnectToMonitor(monitor Monitor)
	GetOS() string
}
