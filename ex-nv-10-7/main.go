package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	c := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			for i := 0; i < 10; i++ {
				c <- i
			}
			wg.Done()
		}(i)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(<-c)
			wg.Done()
		}()
	}

	wg.Wait()
}
