package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var f float64 = 100

	wg.Add(50)

	for i := 0; i < 50; i++ {
		go func() {

			x := math.Sqrt(f)
			f++
			fmt.Println(f, x)
			wg.Done()
		}() //(9.0,&wg)
	}
	wg.Wait()
}
