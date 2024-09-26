package main

/*
func main() {

	c := make(chan int, 5)

	go func(c chan int) {
		for i := 1; i <= 10; i++ {
			fmt.Printf("go-routine #%d func sends data into the channel\n", i)
			c <- i
			fmt.Printf("func go-routine #%d after value is sent into the channel\n", i)
		}
		close(c)
	}(c)
	fmt.Println("making main go-routine sleep so others have time to start")
	time.Sleep(time.Second * 6)
	fmt.Println("receiving values from the channel...")

	//iterating so that all values come out:
	for value := range c {
		fmt.Println("Main go-routine received value:", value)
	}
}*/
