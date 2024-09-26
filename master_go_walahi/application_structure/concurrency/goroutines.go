package main

/*func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 go-routine has started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1 gives i=", i)
	}

	fmt.Println("f1 has ended")
	wg.Done()

}

func f2() {
	fmt.Println("function f2 has started")
	for i := 5; i < 8; i++ {
		fmt.Println("f2 gives i=", i)
	}

	fmt.Println("f2 has ended")
}

func main() {

	//using waitgroups to ensure go-routine is fully executed
	var wg sync.WaitGroup
	wg.Add(1)

	/*fmt.Println("pointless program to do many things...")
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of go-routines", runtime.NumGoroutine())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Architecture", runtime.GOARCH)*/

/*go f1(&wg)
	fmt.Println("go-routines after f1():", runtime.NumGoroutine())

	f2()
	wg.Wait()
	//time.Sleep(time.Second * 2)
	fmt.Println("execution stopped in main") //f1 doesnt execute main func doesnt wait for go-routine func to execute

}*/
