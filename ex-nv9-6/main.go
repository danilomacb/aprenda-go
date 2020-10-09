package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("O sistema operacional é:", runtime.GOOS)
	fmt.Println("A arquitetura é:", runtime.GOARCH)
}
