package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d,cap=%d\n",
		len(s), cap(s))
}

func sliceOps() {
	fmt.Println("Creating slice")
	var s []int

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	printSlice(s2)
	copy(s2, s1)
	printSlice(s1)
	printSlice(s2)
	fmt.Println(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println(s2)
	printSlice(s2)
	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	fmt.Println(s2)
	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)
	//s2 =s2[1:]
	//s2 =s2[1:]
	//for i:=0;i<100;i++{
	//	s2=append(s2,i+3)
	//	printSlice(s2)
	//}
	s2 = add_front(s2, 5)
	fmt.Println(s2)
	printSlice(s2)

	fmt.Println(s1)
	reverse(s1)
	fmt.Println(s1)

	s = []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:3])
	reverse(s[3:])
	reverse(s[:])
	fmt.Println(s)

	data := []string{"one", "", "there"}
	fmt.Printf("%q\n", nonempty2(data))
	fmt.Printf("%q\n", data)
}

func main() {
	sliceOps()
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}

	return out
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func add_front(a []int, d int) []int {
	b := []int{d}
	fmt.Println(a)
	b = append(b, a...)
	return b
}
