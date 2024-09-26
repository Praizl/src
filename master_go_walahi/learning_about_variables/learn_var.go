package main

import (
	"fmt"
)

var weGlobalNow = "Chee"

func main() {
	var name = "aro"
	fmt.Println("my name is", name)

	var age int = 26
	fmt.Println("this is my age", age)

	car, cost := "My car is Mercedes G63 AMG and it cost", 65000000
	fmt.Println(car, cost)
	car, year := "Second drive is Urus", 2022
	fmt.Println(car, year)

	var (
		salary     float64
		MaidenName string
		gender     bool
	)

	fmt.Println(salary, MaidenName, gender)
	var a, b, c int
	fmt.Println(a, b, c)

	//calling the global and local variable within function block
	var localRappers = "lololo"
	fmt.Printf("The global variable is: %s,\nAnd when I call local variable you say \"%v\"\n", weGlobalNow, localRappers)

	display()

	//complex variables
	var w = 6 + 9i
	var v = 9 + 6i
	var p = v + w
	fmt.Println(p)

	//boolean
	d := 2 + 2i
	f := 2 + 3i
	e := 2 + 2i
	same := d == f
	notSame := d == e
	fmt.Println(same, notSame)
	fmt.Printf("d and f are %v, while d and e are %v", same, notSame)
}

func display() {
	fmt.Printf("global variable %v for the w\n", weGlobalNow)
}
