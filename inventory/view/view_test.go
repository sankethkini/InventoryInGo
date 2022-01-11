package view

import (
	"inventory/inventory/commands"
	"reflect"
	"testing"
)

var tests = []struct {
	inp int
	typ interface{}
}{
	{
		inp: 1,
		typ: &commands.Add{},
	},
	{
		inp: 2,
		typ: &commands.Display{},
	},
}

func TestUserInput(t *testing.T) {
	for _, val := range tests {
		UserInput(val.inp, "sanket", 12, 1223, "raw")
		if reflect.TypeOf(curCommand) != reflect.TypeOf(val.typ) {
			t.Error("wrong type")
		}
	}
}
