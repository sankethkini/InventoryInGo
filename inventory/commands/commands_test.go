package commands

import (
	"inventory/inventory/item"
	"reflect"
	"testing"
)

var test = []struct {
	name     string
	price    float64
	quantity int
	typ      string
	exp      interface{}
}{
	{
		name:     "sanket",
		price:    300,
		quantity: 10,
		typ:      "raw",
		exp:      item.RawItem{},
	},
	{
		name:     "sanket",
		price:    300,
		quantity: 10,
		typ:      "imported",
		exp:      item.ImportedItem{},
	},
	{
		name:     "sanket",
		price:    300,
		quantity: 10,
		typ:      "manufactured",
		exp:      item.ManufacturedItem{},
	},
}

func TestAdd(t *testing.T) {
	addcommand := Add{}
	for _, val := range test {
		addcommand.Init(val.name, val.quantity, val.price, val.typ)
		if reflect.TypeOf(addcommand.cur) != reflect.TypeOf(val.exp) {
			t.Error("not of same type")
		}
	}
}
