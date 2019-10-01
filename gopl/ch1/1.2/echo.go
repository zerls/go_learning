package main

import (
	"fmt"
	"os"
	"strings"
)

func demo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
func demo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
func demo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}
func demo4() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}

}

func main() {
	demo4()
}
