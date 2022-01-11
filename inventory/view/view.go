package view

import (
	"fmt"
	"inventory/inventory/commands"
)

var curCommand commands.Command

func GetDetails() {
	var name, typ string
	var quantity int
	var price float64
	fmt.Println("Enter name of the product")
	fmt.Scanf("%s", &name)
	fmt.Println("Enter type of product raw,imported,manufactured")
	fmt.Scanf("%s", &typ)
	fmt.Println("Enter the quantity")
	fmt.Scanf("%d", &quantity)
	fmt.Println("Enter the Price")
	fmt.Scanf("%f", &price)
	UserInput(1, name, quantity, price, typ)
}

func Display(cur []map[string]string) {

	for _, val := range cur {
		fmt.Println("Name :", val["name"], "Price :", val["price"], "Tax :", val["tax"], "Total :", val["total"])
	}
}
func UserInput(curInput int, name string, quantity int, price float64, typ string) {
	if curInput == 1 {
		add := commands.Add{}
		add.Init(name, quantity, price, typ)
		curCommand = &add
		cur := curCommand.Execute()
		Display(cur)
	} else if curInput == 2 {
		curCommand = &commands.Display{}
		curCommand.Execute()
		cur := curCommand.Execute()
		Display(cur)
	} else {
		curCommand = &commands.Exit{}
		cur := curCommand.Execute()
		Display(cur)
	}
}
