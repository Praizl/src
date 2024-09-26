package main

import (
	"fmt"
	//"strings"
)

func main() {

	/*var numbers []int

	a := [...]int{4, 3, 2, 0}
	fmt.Println(numbers, a)
	fmt.Printf("%#v\t", a)
	fmt.Println(len(a))
	a[3] = 1
	for index, value := range a {
		fmt.Println(index, value)
	}
	fmt.Println(strings.Repeat("#", 40))

	//iterating over an array using for C# similar code
	for i := 0; i < len(a); {
		fmt.Println("index:", i, "value:", a[i])
		i++
	}
	b := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(b)
	//keyed arrays
	c := [5]string{4: "Dan",
		0: "Tunji",
		"Doc"}
	fmt.Printf("%#v", c)*/

	//ways to iterate over an array (using for, normal range, as a while loop using range)
	e := [4]int{2, 4, 6, 8}
	for i := 0; i < len(e); i++ {
		fmt.Println(e[i])
	}

	for index, value := range e {
		fmt.Println("index =>", index, "value=>", value)
	}
	j := 0
	for range e {
		fmt.Println(e[j])
		j++
	}
	//multi-dimensional array
	multiArr := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	fmt.Println(multiArr)
	fmt.Println(len(multiArr))
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%v\t", multiArr[i][j])
		}
		fmt.Println("")
	}
}
