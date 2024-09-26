package main

import (
	"fmt"
)

func main() {
	//this shows the different data types and how to use them

	//variables
	//constants
	//string
	//bool

	//array that can contain 6 elements of int type
	var p = [6]int{1, 6, 9, 5, 7}
	fmt.Printf("%T\n", p)

	//slice has dynamic no of elements
	var y = []string{"2", "5"}
	fmt.Printf("%T\n", y)

	//map:group of elements of one type indexed by keys of another type
	var x = map[string]int{
		"Black":     32,
		"Caucasian": 29,
		"Asian":     30,
	}
	fmt.Printf("%T\n", x)
	fmt.Println(x)

	//struct (creates a new user named type which you can assign to a variable)
	type Women struct {
		race string
		age  int
	}
	var Banny Women

	fmt.Println(Banny.age, Banny.race)
}
