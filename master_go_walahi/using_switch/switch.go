package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	language := os.Args[1]
	number, err := strconv.Atoi(os.Args[2])
	_ = err
	switch language {
	case "Python":
		fmt.Println("You are learning python")
	case "Go", "golang":
		fmt.Println("Golang is decent, keep at it")
	case "C++", "javascript", "rubyonrails":
		fmt.Println("Other languages are also good")
	default:
		fmt.Println("Not a programming language, try again!")

	}
	switch true {
	case number%2 == 0:
		fmt.Printf("%v is an Even number\n", number)
	case number%2 != 0:
		fmt.Printf("%v is an Odd number\n", number)
	default:
		fmt.Printf("%v is neither an even or an odd number!\n", number)

	}
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("It's early. flex.")
	case hour < 15:
		fmt.Println("Still have a little time, waste it!")
	case hour < 17:
		fmt.Println("Coming soon,Try Again tommorrow!")
	}
	var value interface{} = "Balablu"
	switch a := value.(type) {
	case string:
		fmt.Println("", a)
	case float64:
		fmt.Println("", a)
	case int64:
		fmt.Println("", a)
	}
}
