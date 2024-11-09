package main

import "fmt"

//2. Multiple Return Values
// func vals() (int, int) {
// 	return 3, 7
// }

// func main() {
// 	a, b := vals()
// 	fmt.Println(a)
// 	fmt.Println(b)

// 	_, c := vals()
// 	fmt.Println(c)
// }

// 3. Variadic Functions
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func main() {
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
