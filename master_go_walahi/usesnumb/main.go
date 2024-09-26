package main

import (
	"fmt"
	"src/mypackages/numbers"
)

func main() {
	var x uint = 23
	y := numbers.Even(x)
	fmt.Printf("%d is an even number:%t\n", x, y)
	fmt.Println(numbers.IsPrime(23), numbers.IsPrime(7))
}
