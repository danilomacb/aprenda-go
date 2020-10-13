package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	q := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

		q <- 0
	}()

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
