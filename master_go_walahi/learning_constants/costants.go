package main

import (
	"fmt"
)

func main() {

	const days = 5
	const secondsperhour = 3600
	hours := 120

	fmt.Printf("number of seconds in a %v days are: %d", days, hours*secondsperhour)

	const a, b = 6, 1
	fmt.Println(a / b) //cant divide by zero

	const p = 5
	const y = 2.2 * p
	fmt.Println(y)

	//using IOTA
	const (
		test1 = iota + 2
		test2
		test3
	)
	fmt.Println(test1, test2, test3)
}
