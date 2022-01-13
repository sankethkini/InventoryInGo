package commands

import (
	"fmt"
	"inventory/inventory/item"
	"os"
)

var items []item.Item

//createmsg creates required return message
func createmsg(msg string) (returnmsg []data) {

	mp := make(data)
	mp["message"] = msg
	returnmsg = append(returnmsg, mp)
	return
}

//alisaing map[string]string
type data = map[string]interface{}

type Command interface {
	Execute() ([]data, error)
}

type add struct {
	cur item.Item
}

func (add *add) Execute() ([]data, error) {
	items = append(items, add.cur)
	msg := createmsg("item added successfuly")
	return msg, nil
}

type display struct {
}

func (display *display) Execute() ([]data, error) {
	var allitems []data
	for _, val := range items {
		cur := make(data)
		name, price, quantity, tax := val.GetDetails()
		cur["name"] = name
		cur["price"] = price
		cur["quantity"] = quantity
		cur["total"] = val.Calc()
		cur["tax"] = tax
		allitems = append(allitems, cur)
	}
	return allitems, nil
}

type exit struct {
}

func (exit *exit) Execute() ([]data, error) {
	fmt.Println("exiting....")
	os.Exit(0)
	return nil, nil
}

func NewAddCommand(name string, quantity int, price float64, typ string) (Command, error) {
	add := add{}
	switch typ {
	case "raw":
		add.cur = item.NewRawItem(name, price, quantity)
	case "imported":
		add.cur = item.NewImportedItem(name, price, quantity)

	case "manufactured":
		add.cur = item.NewManufacturedItem(name, price, quantity)
	default:

		return nil, NotARightTypeErr
	}
	return &add, nil
}

func NewDisplayCommand() (Command, error) {
	display := display{}
	return &display, nil
}

func NewExitCommand() (Command, error) {
	exit := exit{}
	return &exit, nil
}
