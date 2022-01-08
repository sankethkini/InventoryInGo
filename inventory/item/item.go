package item

import "inventory/inventory/constants"

type BaseItem struct {
	Name     string
	Price    float64
	Quantity int
	Tax      float64
}

//every type of item should implement this
type Item interface {
	GetDetails() BaseItem
	Calc() float64
}

//function to check negetive number
func checkneg(cur float64) bool {
	return cur <= 0
}

//item of type raw
type RawItem struct {
	B BaseItem
}

func (item RawItem) GetDetails() BaseItem {
	baseItem := BaseItem{Name: item.B.Name, Price: item.B.Price, Quantity: item.B.Quantity, Tax: item.B.Tax}
	return baseItem
}
func (item RawItem) Calc() float64 {
	if checkneg(item.B.Price) {
		return 0
	}
	total := item.B.Price
	total += total * (constants.RawTax / 100)
	return total
}

//item of type imported
type ImportedItem struct {
	B BaseItem
}

func (item ImportedItem) GetDetails() BaseItem {
	baseItem := BaseItem{Name: item.B.Name, Price: item.B.Price, Quantity: item.B.Quantity, Tax: item.B.Tax}
	return baseItem
}
func (item ImportedItem) Calc() float64 {
	if checkneg(item.B.Price) {
		return 0
	}
	total := item.B.Price
	total += total * (constants.ImportTax / 100)
	if total <= 100 {
		total += constants.Surcharge100
	} else if total <= 200 {
		total += constants.Surcharge200
	} else {
		total += total * (constants.SurchargeMore / 100)
	}
	return total
}

//item of type Manufactured
type ManufacturedItem struct {
	B BaseItem
}

func (item ManufacturedItem) GetDetails() BaseItem {
	baseItem := BaseItem{Name: item.B.Name, Price: item.B.Price, Quantity: item.B.Quantity, Tax: item.B.Tax}
	return baseItem
}
func (item ManufacturedItem) Calc() float64 {
	if checkneg(item.B.Price) {
		return 0
	}
	total := item.B.Price
	total += total * (constants.ManufacturedTax / 100)
	total += total * (constants.ManufacturedExtra / 100)
	return total
}
