package main

import (
	"flag"

	"github.com/sankethkini/InventoryInGo/inventory/view"
)

func main() {
	namePtr := flag.String("name", "unknown", "name of the product")
	pricePtr := flag.Float64("price", 0, "price of the item")
	quantityPtr := flag.Int("quantity", 0, "quantity")
	typePtr := flag.String("type", "raw", "type of the item (raw,manufactured,imported)")
	flag.Parse()
	if *namePtr != "unknown" {
		view.RunCommand(1, *namePtr, *quantityPtr, *pricePtr, *typePtr)
	}
	for {
		view.MenuForUser()
	}
}
