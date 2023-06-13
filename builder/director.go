package builder

type Director struct {
	builder HouseBuilder
}

func (d *Director) SetBuilder(b HouseBuilder) {
	d.builder = b
}

func (d *Director) BuildMommmyHouse() (house House, err error) {
	d.builder.AddChimney()
	d.builder.AddElectricalVoltage(1300)
	d.builder.AddFloor(
		NewFamilySpace().
			AddDoor().
			AddWindow().
			AddRoom("family room").
			AddRoom("bedroom").
			AddRoom("kitchen").
			AddRoom("bathroom").
			AddGarage(normalGarage).Build(),
	)
	d.builder.AddFloor(
		NewBasement().
			AddDoor().
			AddGarage(basementGarage).Build(),
	)

	house, err = d.builder.Build()
	return
}
