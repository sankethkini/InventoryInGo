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
	{Name: "one", Price: 50, Quantity: 12, ExpcRaw: 56.25, ExpcManf: 57.375, ExpcImp: 60},
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

func TestCreateAndDisplay(t *testing.T) {
	var test = []struct {
		Name     string
		Price    float64
		Quantity int
		Typ      string
		Tax      float64
		Total    float64
	}{
		{
			Name:     "apples",
			Price:    400,
			Quantity: 15,
			Typ:      "raw",
			Tax:      12.5,
			Total:    450,
		},
	}

	for _, val := range test {
		raw := NewRawItem(val.Name, val.Price, val.Quantity)
		mt := NewMainItem(&raw)
		gotName, gotPrice, gotQuan, gotTax := mt.GetDetails()
		gotTotal := mt.Calc()
		if gotName != val.Name || gotPrice != val.Price || gotQuan != val.Quantity || gotTax != val.Tax || gotTotal != val.Total {
			t.Errorf("not got as expected")
			t.Errorf("expected %v got %v\n", val.Name, gotName)
			t.Errorf("expected %v got %v\n", val.Price, gotPrice)
			t.Errorf("expected %v got %v\n", val.Quantity, gotQuan)
			t.Errorf("expected %v got %v\n", val.Tax, gotTax)
			t.Errorf("expected %v got %v\n", val.Total, gotTotal)
		}
	}
}
