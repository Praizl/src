package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*fmt.Println("what osArg contains:", os.Args)
	fmt.Println("This shows path", os.Args[0])
	fmt.Println("first argument:", os.Args[1])
	fmt.Println("second argument:", os.Args[2])
	fmt.Println("nth argument", os.Args[3])*/
	argOneinStr, err := strconv.ParseFloat(os.Args[1], 64)
	_ = err
	fmt.Println(argOneinStr)
	fmt.Printf("type for variable is %T value as float after conversion is %v", os.Args[1], argOneinStr)
	x := 0
label2:
	for x < 10 {
		fmt.Println(x)

		if x == 7 {
			x = x + 2
			goto label2
		}
		x++
	}
}
