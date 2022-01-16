package item

import "testing"

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
	for _, val := range tests {

		raw := NewRawItem(val.Name, val.Price, val.Quantity)
		if raw.Calc() != val.ExpcRaw {
			t.Errorf("error in calculating raw product price exp: %v got: %v", val.ExpcRaw, raw.Calc())
		}

		imported := NewImportedItem(val.Name, val.Price, val.Quantity)
		if imported.Calc() != val.ExpcImp {
			t.Errorf("error in calculating imported product price exp: %v got: %v", val.ExpcImp, imported.Calc())
		}

		manf := NewManufacturedItem(val.Name, val.Price, val.Quantity)
		if manf.Calc() != val.ExpcManf {
			t.Errorf("error in calculating manufactured product price exp: %v got: %v", val.ExpcManf, manf.Calc())
		}
	}
}
