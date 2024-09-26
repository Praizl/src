package main

import (
	"fmt"
	"strconv"
)

func main() {

	//int to string
	s := string(69)

	//float to string
	e := fmt.Sprintf("%v", 44.2)
	fmt.Printf("first value %v and the second type %T\n", s, e)
	//string to numbers
	x := "55.3"
	y, err := strconv.ParseFloat(x, 64)
	_ = err
	fmt.Printf("type is %T and value of string is %v\n", y, y)

	//itoA and Atoi
	m, err := strconv.Atoi("26")
	r := strconv.Itoa(26)
	_ = err
	fmt.Println(m)
	fmt.Println(r)
}
