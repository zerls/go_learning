package main

import (
	"fmt"
)

func main() {
	//defer_call()
	a := [3]int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Print(a)
		}
		a[k] = 100 + v
	}
	fmt.Println()
	fmt.Print(a)
	fmt.Println()
	b := []int{1, 2, 3}
	for k, v := range b {
		if k == 0 {
			b[0], b[1] = 100, 200
			fmt.Print(b)
		}
		b[k] = 100 + v
	}
	fmt.Println()
	fmt.Print(b)

}

func defer_call() {

	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()

	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

/*考点：defer
defer 是后进先出。
panic 需要等defer 结束后才会向上传递。 出现panic恐慌时候，会先按照defer的后入先出的顺序执行，最后才会执行panic。

*/
