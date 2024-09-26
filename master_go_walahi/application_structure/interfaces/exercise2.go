package main

import "fmt"

type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {
	var x interface{}
	x = cube{edge: 5}
	fmt.Println(x)
	v := volume(cube{edge: 5}) //v:=volume(x.(cube))
	fmt.Printf("Cube Volume: %v\n", v)

	var empty interface{}
	empty = 7
	fmt.Printf("%T\n", empty)

	empty = 5.8
	fmt.Printf("%T\n", empty)

	empty = []int{}
	fmt.Printf("%T\n", empty)

	empty = append(empty.([]int), 69)
	fmt.Println(empty)
}
