package commands

import (
	"fmt"
	"os"

	"github.com/sankethkini/InventoryInGo/inventory/item"

	"github.com/pkg/errors"
)

var items []*item.MainItem

// createMsg creates required return message.
func createMsg(msg string) (returnmsg []data) {
	mp := make(data)
	mp["message"] = msg
	returnmsg = append(returnmsg, mp)
	return
}

// alisaing map[string]string.
type data = map[string]interface{}

// Command is a interface for all the commands.
type Command interface {
	Execute() ([]data, error)
}

// add command implementation.
type add struct {
	cur *item.MainItem
}

// NewAddCommand is add command's constructor.
func NewAddCommand(name string, quantity int, price float64, typ string) (Command, error) {
	add := new(add)
	switch typ {
	case "raw":
		raw := item.NewRawItem(name, price, quantity)
		add.cur = item.NewMainItem(&raw)

	case "imported":
		imported := item.NewImportedItem(name, price, quantity)
		add.cur = item.NewMainItem(&imported)

	case "manufactured":
		manf := item.NewManufacturedItem(name, price, quantity)
		add.cur = item.NewMainItem(&manf)

	default:

		return nil, errors.Wrap(ErrNotARightType, "cannot create item")
	}

	return add, nil
}

func (add *add) Execute() ([]data, error) {
	items = append(items, add.cur)
	msg := createMsg("item added successfully")
	return msg, nil
}

// display command implementation.
type display struct{}

// NewDisplayCommand is display command's constructor.
func NewDisplayCommand() (Command, error) {
	display := new(display)
	return display, nil
}

func (display *display) Execute() ([]data, error) {
	allItems := make([]data, 0, len(items))

	for _, val := range items {
		cur := make(data)
		name, price, quantity, tax := val.GetDetails()

		cur["name"] = name
		cur["price"] = price
		cur["quantity"] = quantity
		cur["total"] = val.Calc()
		cur["tax"] = tax

		allItems = append(allItems, cur)
	}

	return allItems, nil
}

// exit command implementation.
type exit struct{}

// NewExitCommand is exit command's constructor.
func NewExitCommand() (Command, error) {
	exit := new(exit)
	return exit, nil
}

func (exit *exit) Execute() ([]data, error) {
	fmt.Println("exiting....")
	os.Exit(0)
	return nil, nil
}
