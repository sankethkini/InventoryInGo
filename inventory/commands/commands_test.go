package commands

import (
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

var test = []struct {
	typ      string
	expected error
	wantErr  bool
}{
	{
		typ:      "raw",
		expected: nil,
		wantErr:  false,
	},
	{
		typ:      "manufactured",
		expected: nil,
		wantErr:  false,
	},
	{
		typ:      "imported",
		expected: nil,
		wantErr:  false,
	},
	{
		typ:      "other",
		expected: NotARightTypeErr,
		wantErr:  false,
	},
}

func TestAdd(t *testing.T) {

	for _, val := range test {
		_, err := NewAddCommand("apples", 10, 100, val.typ)
		if val.wantErr && errors.Cause(err).Error() != val.expected.Error() {
			t.Errorf("add command is not correct expected: %v got :%v", val.expected, err)
		}
	}
}

func TestAddAndDisply(t *testing.T) {

	result := make(map[string]interface{})
	result["name"] = "apples"
	result["price"] = float64(800)
	result["quantity"] = 15
	result["total"] = float64(800*0.125 + 800)
	result["tax"] = float64(12.5)

	var alltests = []struct {
		Name     string
		Price    float64
		typ      string
		Quantity int
		wantErr  bool
		res      map[string]interface{}
	}{
		{
			Name:     "apples",
			Price:    800,
			typ:      "raw",
			Quantity: 15,
			wantErr:  false,
			res:      result,
		},
	}

	for _, val := range alltests {
		add, err := NewAddCommand(val.Name, val.Quantity, val.Price, val.typ)

		if !val.wantErr && err != nil {
			t.Errorf("unexpected error occured %v", err)
		}

		_, err = add.Execute()

		if !val.wantErr && err != nil {
			t.Errorf("unexpected error occured %v", err)
		}

		disp, err := NewDisplayCommand()

		if !val.wantErr && err != nil {
			t.Errorf("unexpected error occured %v", err)
		}

		got, err := disp.Execute()

		if !val.wantErr && err != nil {
			t.Errorf("unexpected error occured %v", err)
		}

		if !reflect.DeepEqual(got[0], val.res) {
			t.Error(got, result)
		}

	}
}

func TestExit(t *testing.T) {
	_, err := NewExitCommand()
	if err != nil {
		t.Errorf("expected no error got:% v", err)
	}
}
