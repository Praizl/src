package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func call(y int) {
	fmt.Println(y)

}

func cube(x float64) float64 {
	return math.Pow(x, 3)
}

func f1(n uint) (int, int) {
	factorial := 1
	sum := 0

	for i := 1; i < int(n); i++ {
		factorial *= i
		sum += i
	}

	return factorial, sum
}

func myfunc(n string) int {
	arg, err := strconv.Atoi(n)
	if err != nil {
		fmt.Printf("cannot convert string %q to int value", n)
		os.Exit(1)
	}
	arg2, _ := strconv.Atoi(n + n)
	arg3, _ := strconv.Atoi(n + n + n)

	return arg + arg2 + arg3

}

func main() {
	z := call
	z(1)
	fmt.Printf("%T\n", z)

	cube(3)
	fmt.Println(cube(3))

	fmt.Println(f1(5))

	fmt.Println(myfunc("1"))
}
