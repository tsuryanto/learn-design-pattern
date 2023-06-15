package document

import (
	"fmt"
	"learndesignpattern/prototype"
)

type IdCard struct {
	Name    string
	Company string
	barcode string
}

func NewIdCard(name string, company string) IdCard {
	return IdCard{
		Name:    "Taufiq",
		Company: "Halo Kitchen",
		barcode: "X334rg",
	}
}

func (i IdCard) GetBarcode() string {
	return i.barcode
}

func (i IdCard) Print() {
	fmt.Println("ID Card + ", i.Name)
}

func (i IdCard) Clone() prototype.Document {
	return &IdCard{
		Name:    i.Name,
		Company: i.Company,
		barcode: i.barcode,
	}
}
