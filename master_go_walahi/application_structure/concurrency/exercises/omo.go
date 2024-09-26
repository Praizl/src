package main

/*

func main() {
	var wg sync.WaitGroup

	balance := 100
	var n int
	wg.Add(200)
	var m sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(time.Second * 5)
			m.Lock()
			balance += n
			m.Unlock()

			wg.Done()
		}()
		go func() {
			time.Sleep(time.Second * 5)
			m.Lock()
			defer m.Unlock()
			balance -= n

			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}
*/
