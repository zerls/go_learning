package spinner

import (
	"fmt"
	"time"
)

func main() {
	before := time.Now()
	go spinner(100 * time.Millisecond)
	const n = 25
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d)= %d \n", n, fibN)
	afer := time.Now()
	fmt.Println(afer.Sub(before))
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
