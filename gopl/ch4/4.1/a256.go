package main

import (
	"crypto/sha256"
	"fmt"
)

//var(
//	me =flag.String("m","sha256","select sha256/384/512")
//	val =flag.String("v","","need sha value")
//)

func main() {
	//flag.Parse()
	//r := [...]int{99: 15,1,-1}
	//fmt.Println(r)

	//4.2
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	var c3 [32]byte
	for i, _ := range c1 {
		c3[i] = c1[i] ^ c2[i]
	}
	//fmt.Printf("%x\n%x\n%t\n%T\n",c1,c2,c1==c2,c1)
	//var v1,v2 int
	//fmt.Scanf("%d %d",&v1,&v2)
	//fmt.Printf("%d + %d = %d\n",v1,v2,v1+v2)
	num := 0
	for _, c := range c3 {
		num += PopCount(uint64(c))
		fmt.Printf("%b %d\n", c, PopCount(uint64(c)))
	}
	//fmt.Println()
	fmt.Println(num)

	//4.2
	//x :=[]byte(*val)
	//fmt.Println(*val)
	//switch *me {
	//case "sha256":
	//	c3 :=sha256.Sum256([]byte(x))
	//	fmt.Printf("[%s] %x\n",*me,c3)
	//case "sha384":
	//	c3 :=sha512.Sum384([]byte(x))
	//	fmt.Printf("[%s] %x\n",*me,c3)
	//case "sha512":
	//	c3:=sha512.Sum512([]byte(x))
	//	fmt.Printf("[%s] %x\n",*me,c3)
	//default:
	//	fmt.Printf("Please select -m right arg")
	//}

}

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Printf("%b %d\n", byte(i&1), pc[i])
	}

}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
