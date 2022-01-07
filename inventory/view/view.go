package view

import (
	"fmt"

	"../commands"
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

//comment
func UserInput(curInput int, name string, quantity int, price float64, typ string) {
	if curInput == 1 {
		add := commands.Add{}
		add.Init(name, quantity, price, typ)
		curCommand = &add
		curCommand.Execute()
	} else if curInput == 2 {
		curCommand = &commands.Display{}
		curCommand.Execute()
	} else {
		curCommand = &commands.Exit{}
		curCommand.Execute()
	}
}
