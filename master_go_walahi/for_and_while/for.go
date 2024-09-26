package main

import "fmt"

func main() {
	for a := 0; a < 13; a++ {
		fmt.Println(a)
	}
	b := 9
	for b >= 5 {
		fmt.Println(b)
		b--
	}
	//continue
	for i := 0; i < 25; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	for j := 5; j <= 30; j++ {
		if j%5 != 0 || j%2 != 0 {
			continue
		}
		fmt.Println(j)

	}
	//break
	for i := 5; i < 50; i += 5 {
		if i%10 == 0 {
			i = i + 1
			continue
		} else {
			fmt.Println(i)
		}

	}
	count := 0
	for i := 0; i < 120; i++ {
		if i%13 == 0 {
			fmt.Printf("%v is divisible by 13\n ", i)
			count++
		}
		if count == 5 {
			break
		}
	}
	fmt.Println(count)

}
