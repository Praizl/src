package main

import "fmt"

func changeValue(quantity int, price float64, name string, sold bool) {
	quantity = 4
	price = 399.99
	name = "Solomon Grundy"
	sold = false
}

func changeByPointer(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 12
	*price = 699.99
	*name = "Chinua Achebe"
	*sold = false

}

type product struct {
	price       float64
	productName string
}

func changeStruct(gift product) {
	gift = product{
		price:       120,
		productName: "Mens Briefs",
	}

}

func changeStructByPtr(gift *product) {
	gift.price = 120
	gift.productName = "Mens Briefs"
}

func main() {
	r := 3
	pr := &r

	fmt.Println(r, pr, *pr)
	prpr := &pr
	fmt.Println(prpr, **prpr)

	**prpr = 69

	fmt.Println(r, pr, *pr, prpr)
	fmt.Println(pr == *prpr)

	s := 69
	newPr := &s
	fmt.Println(*newPr, *pr, newPr == pr) //different addresses so not equal.

	quantity := 3
	price := 499.99
	name := "William Shakespeare "
	sold := true

	changeValue(quantity, price, name, sold)
	fmt.Println(quantity, price, name, sold) //it didnt change
	changeByPointer(&quantity, &price, &name, &sold)
	fmt.Println(quantity, name, sold, price) //changes (because we passed pointers of the variables(&v) to the changebypointer fn  )

	gift := product{
		price:       750.50,
		productName: "PS.5",
	}
	changeStruct(gift)
	fmt.Println(gift) //UNCHANGED
	changeStructByPtr(&gift)
	fmt.Println(gift) //changed
}
