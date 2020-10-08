package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go f1()
	go f2()

	wg.Wait()
}

func f1() {
	fmt.Println("Function 1")
	wg.Done()
}

func f2() {
	fmt.Println("Function 2")
	wg.Done()
}