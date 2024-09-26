package main

import (
	"fmt"
	"math"
)

func main() {
	x := 3.1
	result := math.Pow(x, 8)
	fmt.Println("x to the power of 8 is", result)
	//map try
	var Simp = map[string]int{
		"Amala": 5,
		"Ewedu": 12,
		"Beef":  15,
	}
	fmt.Println(Simp)
	//struct
	type Today struct {
		date int
		day  string
	}
	var bruh Today
	bruh.date = 2
	bruh.day = "Sunday"
	fmt.Println(bruh.date, bruh.day)
	fmt.Printf("%T", bruh.date)
	//array
	r := [7]int{2, 8, 6, 0, 7, 1, 4}
	fmt.Println(r)
	//slice
	s := []int{5, 7, 8, 2, 4, 0, 8}
	fmt.Println(s)

	//pointer
	var ss int = 3
	pointr := &ss
	fmt.Println(pointr)

	fmt.Printf("%T", f)

}
func f() {

}
