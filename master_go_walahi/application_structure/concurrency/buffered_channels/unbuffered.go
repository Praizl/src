package main

/*
func main() {
	c1 := make(chan int)
	c2 := make(chan int, 3)
	_ = c2
	go func(c chan int) {
		fmt.Println("go-routine func sends data into the channel")
		c <- 10
		fmt.Println("func go-routine after value is sent into the channel")
	}(c1)

	fmt.Println("making go-routine sleep for 6 secs")
	time.Sleep(time.Minute / 10)
	fmt.Println("main go-routine receives data after 6s")
	d := <-c1
	fmt.Println("data received :", d)

}*/
