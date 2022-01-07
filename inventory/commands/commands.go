package commands

import (
	"fmt"
	"os"

	"../item"
)

var items []item.Item

type Command interface {
	Execute()
}

type Add struct {
	cur item.Item
}

//comment
func (add *Add) Init(name string, quantity int, price float64, typ string) {
	switch typ {
	case "raw":
		raw := item.Item{Name: name, Price: price, Quantity: quantity, Tax: 12.5}
		raw.Calc = item.RawCalc
		add.cur = raw
		break
	case "imported":
		imported := item.Item{Name: name, Price: price, Quantity: quantity, Tax: 10}
		imported.Calc = item.ImportedCalc
		add.cur = imported
		break
	case "manufactured":
		manf := item.Item{Name: name, Price: price, Quantity: quantity, Tax: 12.5}
		manf.Calc = item.ManufacturedCalc
		add.cur = manf
		break
	default:
		fmt.Println("not a proper type")
		os.Exit(1)
		break
	}
}

func (add *Add) Execute() {
	items = append(items, add.cur)
}

type Display struct {
}

func (display *Display) Execute() {
	for _, val := range items {
		fmt.Println("Name :", val.Name, "Price :", val.Price, "Tax :", val.Tax, "Total price:", val.Calc(val.Price))
	}
}

type Exit struct {
}

func (exit *Exit) Execute() {
	fmt.Println("exiting....")
}
