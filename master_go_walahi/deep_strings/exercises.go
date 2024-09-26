package main

import (
	"fmt"
)

func main() {
	s := "你好 Go!"
	t := []rune(s)
	fmt.Printf("%#v\n", t)

	for index, value := range t {
		fmt.Printf("%d   %c\n", index, value)
	}

	u := "你好 Go!"
	v := []byte(u)
	fmt.Printf("%#v\n", v)
	for index, value := range v {
		fmt.Printf("%v>>>>>> %v\n", index, value)
	}
	s1 := "Go is cool!"
	x := s1[0]
	fmt.Println(x)

	// printing the number of runes in the string
	fmt.Println(len(s1))

	s2 := "țară means country in Romanian"
	for i := 0; i < len(s2); i++ {
		fmt.Println(s2[i])

	}
	for byteIndex, runE := range s2 {
		fmt.Printf("%d>>>>>>>%c\n", byteIndex, runE)
	}
}
