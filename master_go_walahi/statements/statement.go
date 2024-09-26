package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	//normal if
	age, err := strconv.Atoi(os.Args[1])
	if age < 30 && age > 1 && err == nil {
		fmt.Printf("You're still young! you have at least %v years\n", 30-age)
	} else if age > 30 && err == nil {
		fmt.Println("you don dey old bruh")
	} else if err != nil {
		fmt.Println("invalid age!")
	} else if age < 1 {
		fmt.Println("invalid age")
	}

	//using simple statement
	if age, err := strconv.Atoi(os.Args[2]); age > 1 && err == nil {
		fmt.Println("no error, age is:", age)
	} else {
		fmt.Println(err)
	}

	//example using args and len
	if age := os.Args; len(age) != 4 {
		fmt.Println("Please input your age!")
	} else if age, err := strconv.Atoi(os.Args[3]); err == nil {
		fmt.Println("your age is:", age)

	} else {
		fmt.Println("invalid age")
	}
}
