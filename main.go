package main

import (
	"flag"
	"fmt"

	"../inventory/inventory/view"
)

func main() {
	//review
	namePtr := flag.String("name", "unknown", "name of the product")
	pricePtr := flag.Float64("price", 0, "price of the item")
	quantityPtr := flag.Int("quantity", 0, "quantity")
	typePtr := flag.String("type", "raw", "type of the item (raw,manufactured,imported)")
	flag.Parse()
	view.UserInput(1, *namePtr, *quantityPtr, *pricePtr, *typePtr)
	for {
		fmt.Println("1. add more items 2. display all 3. exit")
		var cur int
		fmt.Scanf("%d", &cur)
		if cur == 1 {
			view.GetDetails()
		} else {
			view.UserInput(cur, "", 0, 0, "")
		}
	}

}
