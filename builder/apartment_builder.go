package builder

import (
	"errors"
	"fmt"
)

type ApartementBuilder struct {
	House
}

func (a *ApartementBuilder) AddFloor(floor Floor) {
	if len(a.Floors) > 1 {
		floor.AddStairs()
	}
	a.Floors = append(a.Floors, floor)
}

func (a *ApartementBuilder) AddChimney() {
	a.Chimneys++
}

func (a *ApartementBuilder) AddElectricalVoltage(total uint) {
	a.ElectricalVoltage += total
}

func (a *ApartementBuilder) Build() (output House, err error) {
	if len(a.Floors) < 5 {
		err = errors.New("minimum floor number is 5")
		return
	}

	if a.Chimneys < 3 {
		err = errors.New("minimum chimneys is 3")
		return
	}

	if a.ElectricalVoltage/uint(len(a.Floors)) < 1000 {
		err = fmt.Errorf("please add more electrical voltage. Total Recomandation is %d", len(a.Floors)*1000)
		return
	}

	return a.House, nil
}
