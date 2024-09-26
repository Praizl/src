package main

import "fmt"

func main() {
	mySlice := []float64{1.2, 5.6}

	mySlice[1] = float64(6)

	a := float64(10)
	mySlice[0] = a

	mySlice[1] = 10.10

	mySlice = append(mySlice, a)

	fmt.Println(mySlice)

	//Ex.8
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	ny := years[0:3]
	newYears := make([]int, 0)
	newYears = append(ny, years[len(years)-3:]...)
	fmt.Println(newYears)

	//exercise 5
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	sums := 0
	for _, valueNums := range nums[2 : len(nums)-2] {

		sums += valueNums
		fmt.Printf("value:%d,\n sums:%d\n", valueNums, sums)
	}

}
