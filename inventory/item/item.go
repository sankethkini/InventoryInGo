package item

import "inventory/inventory/constants"

type baseItem struct {
	name     string
	price    float64
	quantity int
	tax      float64
}

//every type of item should implement this
type Item interface {
	GetDetails() (string, float64, int, float64)
	Calc() float64
}

func (b *baseItem) SetDetails(name string, price float64, quantity int) {
	b.name = name
	b.price = price
	b.quantity = quantity
}

func (b *baseItem) GetDetails() (name string, price float64, quantity int, tax float64) {
	name = b.name
	price = b.price
	quantity = b.quantity
	tax = b.tax
	return
}

//function to check negetive number
func checkneg(cur float64) bool {
	return cur <= 0
}

//rawItem is struct for item of type raw
type rawItem struct {
	baseItem
}

func (item rawItem) Calc() float64 {
	if checkneg(item.price) {
		return 0
	}
	total := item.price
	total += total * (constants.RawTax / 100)
	return total
}

//importedItem is struct for imported items
type importedItem struct {
	baseItem
}

func (item importedItem) Calc() float64 {
	if checkneg(item.price) {
		return 0
	}
	total := item.price
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

//ManufacturedItem is struct for manufactured items
type manufacturedItem struct {
	baseItem
}

func (item manufacturedItem) Calc() float64 {
	if checkneg(item.price) {
		return 0
	}
	total := item.price
	total += total * (constants.ManufacturedTax / 100)
	total += total * (constants.ManufacturedExtra / 100)
	return total
}

func NewManufacturedItem(name string, price float64, quantity int) *manufacturedItem {
	cur := manufacturedItem{}
	cur.tax = constants.ManufacturedTax
	cur.SetDetails(name, price, quantity)
	return &cur
}

func NewImportedItem(name string, price float64, quantity int) *importedItem {
	cur := importedItem{}
	cur.tax = constants.ImportTax
	cur.SetDetails(name, price, quantity)
	return &cur
}

func NewRawItem(name string, price float64, quantity int) *rawItem {
	cur := rawItem{}
	cur.tax = constants.RawTax
	cur.SetDetails(name, price, quantity)
	return &cur
}
