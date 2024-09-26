package main

import (
	"fmt"
)

func main() {
	var name = "Aro"
	fmt.Println("hello, VSCode. I am", name)

	a := 6
	b := 9
	fmt.Println("Sum:", a+b, "Average:", (a+b)/2)

	fmt.Printf("your age is %d\n", 21)

	fmt.Printf("this is how to use \"double quotes\" in a string\n")
	fmt.Printf("x is %d, y is %d\n", 5, 6)

	calc := "circumference of a circle"
	radius := 3
	pi := 3.142

	fmt.Printf("the %s with a radius of %d taking radius as %f is: %.2v\n",
		calc, radius, pi, 2*pi*float64(radius))

	closed := true
	temperature := 25
	quotedstr := "Swiss army knife"

	fmt.Printf("When the bool door is %t, type is %T, to covert to binary is %b, hexadecimal is %x, when unknown use %q percentv ",
		closed, temperature, temperature, temperature, quotedstr)
}
