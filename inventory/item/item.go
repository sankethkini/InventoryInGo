package item

import "inventory/inventory/constants"

//BaseItem contains all the data required for a item
type BaseItem struct {
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

func (b *BaseItem) SetDetails(name string, price float64, quantity int, tax float64) {
	b.name = name
	b.price = price
	b.quantity = quantity
	b.tax = tax
}

func (b *BaseItem) GetDetails() (name string, price float64, quantity int, tax float64) {
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
	BaseItem
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
	BaseItem
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
	BaseItem
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

//NewManufacturedItem is Manufactured item's constructor
func NewManufacturedItem(name string, price float64, quantity int) Item {
	cur := manufacturedItem{}
	cur.SetDetails(name, price, quantity, constants.ManufacturedTax)
	return &cur
}

//NewImportedItem is Imported item's constructor
func NewImportedItem(name string, price float64, quantity int) Item {
	cur := importedItem{}
	cur.SetDetails(name, price, quantity, constants.ImportTax)
	return &cur
}

//NewRawItem is Raw item's constructor
func NewRawItem(name string, price float64, quantity int) Item {
	cur := rawItem{}
	cur.SetDetails(name, price, quantity, constants.RawTax)
	return &cur
}

//MainItem covers all the possible items
type MainItem struct {
	I *Item
}

func (m MainItem) Calc() float64 {
	cur := *m.I
	return cur.Calc()
}

func (m MainItem) GetDetails() (name string, price float64, quantity int, tax float64) {
	cur := *m.I
	return cur.GetDetails()
}

//NewMainItem is mainitem's constructors
func NewMainItem(i *Item) *MainItem {
	return &MainItem{I: i}
}
