package main

/*import "fmt"

func{
	var slice1 []string
	fmt.Println(slice1)

	//different ways to create a slice
	newSlice := []string{"normal way", "using make", "creating new type then assigning"}
	fmt.Println(newSlice)

	usingMake := make([]int, 4, 9)
	fmt.Println(len(usingMake))
	fmt.Printf("%#v\n", usingMake)

	type newType []int
	nutype := newType{1, 2, 3, 4}
	fmt.Println(nutype)

	nutype[3] = 5
	nunu := nutype[3]
	fmt.Println(nunu)
	fmt.Println(nutype)

	//iterating over a slice
	for index, value := range nutype {
		fmt.Println("index:", index, "value", value)
	}

	//using append
	nutype = append(nutype, 69, 04, 3, 10)
	fmt.Println(nutype)

	//appending slice to slice
	slice2 := newType{12, 34, 56, 78}
	fmt.Println(slice2)
	slice2slice := append(nutype, slice2...)
	fmt.Println(slice2slice)

	//using copy
	slice3 := make(newType, len(slice2))
	slice4 := copy(slice3, slice2)
	_ = slice4
	fmt.Println(slice2, slice3)

	//comparing
	var equal bool = true
	for index, value2 := range slice2 {
		if value2 != slice2slice[index] {
			equal = false
			break
		}
	}
	if len(slice2) != len(slice2slice) {
		equal = false
	}
	if equal {
		fmt.Println("The slices are equal")
	} else {
		fmt.Println("The slices are not equal")
	}
}*/
