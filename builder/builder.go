package builder

type GarageType string

const (
	normalGarage   GarageType = "normal"
	basementGarage GarageType = "basement"
)

type HouseBuilder interface {
	AddFloor(floor Floor)
	AddChimney()
	AddElectricalVoltage(total uint)
	Build() (House, error)
}

type FloorBuilder interface {
	AddDoor()
	AddWindow()
	AddRoom(roomName string)
	AddGarage(garageType GarageType)
	Build() Floor
}
