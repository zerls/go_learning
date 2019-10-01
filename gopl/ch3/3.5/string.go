package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

/*
0xxxxxx



*/

func main() {
	//TODO 3.10 3.11 3.12

	s := "Hello, 世界"
	//
	//fmt.Println(len(s))
	//fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	//n:=0
	//for _,_ =range s{
	//	n++
	//}
	//n=0
	//for range s{
	//	n++
	//}
	fmt.Println(utf8.RuneCountInString(s))

	// "program" in Japanese katakana
	s = "プログラム"
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))      // "プログラム"
	fmt.Println(string(65))     // "A", not "65"
	fmt.Println(string(0x4eac)) // "京"

	fmt.Println(string(1234567)) // "�"

	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	fmt.Println(strconv.FormatInt(int64(x), 2))
	x = 8
	s = fmt.Sprintf("x=%x", x)
	fmt.Println(s)

	a, err := strconv.Atoi("123")             // x is an int
	b, err := strconv.ParseInt("123", 10, 64) //base 10, up to 64 bits
	fmt.Println(a, b, err)
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contain(s, substr string) bool {
	for i := 0; i < len(substr); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
