package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	const gr = 100

	var wg sync.WaitGroup

	wg.Add(gr * 2)
	//1.creating mutexes
	var m sync.Mutex
	n := 0

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			n++
			m.Unlock()
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			defer m.Unlock()
			n--
			wg.Done()

		}()
	}
	wg.Wait()

	fmt.Println(n)
}
