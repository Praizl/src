package main

import (
	"fmt"
	"time"
)

func main() {
	started := time.Now().UnixNano() / 1000000
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 5)
		c1 <- "Ansem>"
	}()

	go func() {
		time.Sleep(time.Second * 5)
		c2 <- "Bitboy"

	}()

	time.Sleep(time.Second * 6)
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received c1 value:", msg1)

		case msg2 := <-c2:
			fmt.Println("Received value fromm c2:", msg2)

		default:
			fmt.Println("No activity")
		}
	}

	ended := time.Now().UnixNano() / 1000000
	time := ended - started
	fmt.Println("time:", time)
}
