package builder

type Floor struct {
	doors   uint
	rooms   []string
	windows uint
	garage  []GarageType
	stairs  int
}

func NewFloor() *Floor {
	return &Floor{}
}

type FamilySpace struct {
	Floor
}

func NewFamilySpace() *Floor {
	fam := FamilySpace{}
	return &fam.Floor
}

type Basement struct {
	Floor
}

func NewBasement() *Floor {
	fam := Basement{}
	return &fam.Floor
}

func (f *Floor) AddDoor() *Floor {
	f.doors++
	return f
}

func (f *Floor) AddWindow() *Floor {
	f.windows++
	return f
}

func (f *Floor) AddRoom(name string) *Floor {
	f.rooms = append(f.rooms, name)
	return f
}

func (f *Floor) AddGarage(garageType GarageType) *Floor {
	f.garage = append(f.garage, garageType)
	return f
}

func (f *Floor) AddStairs() *Floor {
	f.stairs++
	return f
}

func (f *Floor) Build() Floor {
	return *f
}
