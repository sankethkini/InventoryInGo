package commands

import (
	"testing"
)

var test = []struct {
	typ      string
	expected error
}{
	{
		typ:      "raw",
		expected: nil,
	},
	{
		typ:      "manufactured",
		expected: nil,
	},
	{
		typ:      "imported",
		expected: nil,
	},
	{
		typ:      "other",
		expected: NotARightTypeErr,
	},
}

func TestAdd(t *testing.T) {

	for _, val := range test {
		_, err := NewAddCommand("apples", 10, 100, val.typ)
		if err != val.expected {
			t.Errorf("add command is not correct expected: %v got :%v", val.expected, err)
		}
	}
}
