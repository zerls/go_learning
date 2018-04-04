package main

import "fmt"

func main() {

	fmt.Println("Hello, World!")

	// 多维数组
	var a = [3][2]int{
		{0, 0},
		{1, 2},
		{2, 4}, //只想说这里为什么非要加逗号，反人类的逗号，但是美观一点
	}
	var n1, n2 int
	for n1 = 0; n1 < 3; n1++ {
		for n2 = 0; n2 < 2; n2++ {
			fmt.Printf("a[%d][%d] = %d\n", n1, n2, a[n1][n2])
		}
	}
}
