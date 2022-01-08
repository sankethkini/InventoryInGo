package item

import (
	"testing"
)

var tests = []struct {
	Name     string
	Price    float64
	Quantity int
	ExpcRaw  float64
	ExpcImp  float64
	ExpcManf float64
}{
	{Name: "one", Price: 400, Quantity: 12, ExpcRaw: 450, ExpcManf: 459, ExpcImp: 462},
	{Name: "two", Price: 0, Quantity: 12, ExpcRaw: 0, ExpcManf: 0, ExpcImp: 0},
	{Name: "", Price: -23, Quantity: 12, ExpcRaw: 0, ExpcManf: 0, ExpcImp: 0},
	{Name: "", Price: 100, Quantity: 12, ExpcRaw: 112.5, ExpcManf: 114.75, ExpcImp: 120},
}

func TestItemCalc(t *testing.T) {
	for i, val := range tests {
		b := BaseItem{Name: val.Name, Price: val.Price, Quantity: val.Quantity}
		raw := RawItem{B: b}
		imported := ImportedItem{B: b}
		manf := ManufacturedItem{B: b}
		if raw.Calc() != val.ExpcRaw {
			t.Error("error in calculating raw product price", i)
		}
		if imported.Calc() != val.ExpcImp {
			t.Error("error in calculating imported product price", i, imported.Calc())
		}
		if manf.Calc() != val.ExpcManf {
			t.Error("error in calculating manufactured product price", i)
		}
	}
}
