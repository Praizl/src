package main

import (
	"fmt"
	"math"
)

func main() {
	//these are the types of arithmectic, logical and comparison operators

	//+,-,/,*

	//modulus
	x := 2
	y := 3
	r := 9 % y
	fmt.Println(r)
	_ = x

	//increment and decrement
	a := 5
	b := 4
	a += b
	fmt.Println(a)

	b -= a
	fmt.Println(b)

	b *= a
	fmt.Println(b)

	b /= -9
	fmt.Println(b)

	a %= b
	fmt.Println(a)

	a, z := 2, 8
	z++
	z += a
	a++
	fmt.Println(a, z)

	//comparison operators
	zuzu := 2
	zazu := 23
	fmt.Println(zuzu < zazu, zazu >= zuzu, zuzu == zazu, zazu != zuzu)
	//logical operators
	fmt.Println(zuzu != zazu && zazu > 20, zazu == zuzu || zazu > 20)

	//uint8 overflow
	m := uint8(255)
	m += 1
	fmt.Println(m)
	yy := int8(127)

	yy += 1
	fmt.Printf("new value will be %v, type %T\n", yy, yy)

	nut := float64(math.MaxFloat64)
	fmt.Println(nut)
	nut = nut * 2
	fmt.Println(nut)

}
