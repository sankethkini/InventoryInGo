package view

import (
	"fmt"
	"inventory/inventory/commands"
)

//getDeatils function helps to get details from user
func getDetails() {
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

	RunCommand(1, name, quantity, price, typ)
}

//display function displays the result of operation to user
func display(cur []map[string]interface{}) {

	fmt.Printf("------------------------\n")
	for _, val := range cur {
		for key, val := range val {
			fmt.Printf("%v ----> %v\n", key, val)
		}
		fmt.Printf("------------------------\n")
	}
}

//executeCommand function executes command selected by user
func executeCommand(cmd commands.Command) {
	res, err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
	display(res)
}

//RunCommand function runs the right command for perticular input
func RunCommand(curInput int, name string, quantity int, price float64, typ string) {

	var curCommand commands.Command
	var err error
	if curInput == 1 {
		curCommand, err = commands.NewAddCommand(name, quantity, price, typ)
	} else if curInput == 2 {
		curCommand, err = commands.NewDisplayCommand()
	} else {
		curCommand, err = commands.NewExitCommand()
	}
	if err != nil {
		fmt.Println(err)

	} else {
		executeCommand(curCommand)
	}

}

//MenuForUser displays the menu for user
func MenuForUser() {
	fmt.Println("1. add more items 2. display all 3. exit")
	var cur int
	fmt.Scanf("%d", &cur)
	if cur == 1 {
		getDetails()
	} else {
		RunCommand(cur, "", 0, 0, "")
	}
}
