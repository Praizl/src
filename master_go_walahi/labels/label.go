package main

import "fmt"

func main() {
	var superstores = [5]string{"Justrite", "Shoprite", "GroceryBazaar", "MarketSquare", "Spar"}
	var inPh = [2]string{"MarketSquare", "Spar"}

outline:
	for index, stores := range superstores {

		for _, inPh := range inPh {
			if stores == inPh {
				fmt.Printf("The store %q is at index %d\n", inPh, index)
				continue outline
			}
		}

	}
	fmt.Println("next step")
	x := 1
label69:
	for x < 5 {
		fmt.Println(x)
		if x == 2 {
			x = x + 2
			goto label69
		}
		x++
	}
}
