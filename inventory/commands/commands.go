package commands

import (
	"fmt"
	"inventory/inventory/constants"
	"inventory/inventory/item"
	"os"
	"strconv"
)

var items []item.Item

//alisaing map[string]string
type data = map[string]string

type Command interface {
	Execute() []data
}

type Add struct {
	cur item.Item
}

func (add *Add) Init(name string, quantity int, price float64, typ string) {
	baseItem := item.BaseItem{Name: name, Quantity: quantity, Price: price}
	switch typ {
	case "raw":
		baseItem.Tax = constants.RawTax
		add.cur = item.RawItem{B: baseItem}

	case "imported":
		baseItem.Tax = constants.ImportTax
		add.cur = item.ImportedItem{B: baseItem}

	case "manufactured":
		baseItem.Tax = constants.ManufacturedTax
		add.cur = item.ManufacturedItem{B: baseItem}

	default:
		fmt.Println("not a proper type")
		os.Exit(1)

	}
}

func (add *Add) Execute() []data {
	items = append(items, add.cur)
	return nil
}

type Display struct {
}

func (display *Display) Execute() []data {
	var allitems []data
	for _, val := range items {
		cur := make(data)
		cur["name"] = val.GetDetails().Name
		cur["total"] = strconv.FormatFloat(val.Calc(), 'f', 3, 64)
		cur["tax"] = strconv.FormatFloat(val.GetDetails().Tax, 'f', 3, 64)
		cur["price"] = strconv.FormatFloat(val.GetDetails().Price, 'f', 3, 64)
		allitems = append(allitems, cur)
	}
	return allitems
}

type Exit struct {
}

func (exit *Exit) Execute() []data {
	fmt.Println("exiting....")
	os.Exit(0)
	return nil
}
