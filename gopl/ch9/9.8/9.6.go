package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
