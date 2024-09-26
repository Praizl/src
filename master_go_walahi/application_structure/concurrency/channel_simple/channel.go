package main

import (
	"fmt"
	"strings"
)

func f1(n int, ch chan int) {
	ch <- n //value in n sent to channel
}

func factorial(n int, channel chan int) {
	f := 2
	for i := 2; i <= n; i++ {
		f *= i
	}
	channel <- f
}

func main() {
	ch := make(chan int) //2dim channel
	//ch:=make(<-chan int)   //to receive int value to channel
	//ch:=make(chan<-int)	 //to send to

	go f1(10, ch)
	n := <-ch //n receives value in channel
	fmt.Println(n)

	channel := make(chan int)
	defer close(channel)

	go factorial(5, channel)
	fact := <-channel
	fmt.Println(fact)

	fmt.Println(strings.Repeat("<>", 30))

	for i := 1; i <= 20; i++ {
		go factorial(i, channel)
		fact := <-channel
		fmt.Println(fact)
	}

	fmt.Println(strings.Repeat("<>", 30))

	for i := 5; i <= 15; i++ { //using anon funcs
		go func(n int, channel chan int) {
			f := 1
			for i := 2; i <= n; i++ {
				f *= i
			}
			channel <- f
		}(i, ch)
		fact = <-ch
		fmt.Println(i, "factorial is", fact)
	}

}
