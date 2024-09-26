package main

import "fmt"

func main() {
	age := 15
	var available bool
	if age >= 18 && age <= 60 && available {
		fmt.Println("Order arriving shortly. NB: Smoking damages your lungs")
	} else if !available {
		fmt.Println("Cigars out of stock!")
	} else if age < 18 {
		fmt.Println("Too young to smoke. Have a candy instead!")
	} else {
		fmt.Println("Too old, stay alive for your grandchildren!")
	}
}
