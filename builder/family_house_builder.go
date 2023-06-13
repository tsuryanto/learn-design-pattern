package builder

import (
	"errors"
	"fmt"
)

type FamilyHouseBuilder struct {
	House
}

func (a *FamilyHouseBuilder) AddFloor(floor Floor) {
	if len(a.Floors) > 1 {
		floor.AddStairs()
	}
	a.Floors = append(a.Floors, floor)
}

func (a *FamilyHouseBuilder) AddChimney() {
	a.Chimneys++
}

func (a *FamilyHouseBuilder) AddElectricalVoltage(total uint) {
	a.ElectricalVoltage += total
}

func (a *FamilyHouseBuilder) Build() (output House, err error) {
	if len(a.Floors) > 2 {
		err = errors.New("maximum floor number is 2")
		return
	}

	if a.Chimneys > 2 {
		err = errors.New("maximum chimneys is 2")
		return
	}

	if a.ElectricalVoltage < 1300 {
		err = fmt.Errorf("please add more electrical voltage. Total Recomandation is %d", 1300)
		return
	}

	return a.House, nil
}
