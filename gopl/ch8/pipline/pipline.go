package main

import "fmt"

//func main(){
//	naturals :=make(chan int)
//	squares :=make(chan int)
//
//	//Counter
//	go func() {
//		for x:=0;x<10;x++{
//			naturals <-x
//		}
//		close(naturals)
//	}()
//
//	//Squarer
//	go func() {
//		for x:=range naturals{
//			squares<-x*x
//		}
//		close(squares)
//	}()
//
//	//go func() {
//	//	for{
//	//		x,ok :=<-naturals
//	//		if !ok{
//	//			break // channel was closed and drained
//	//		}
//	//		squares <- x * x
//	//	}
//	//	close(squares)
//	//}()
//
//	//Printer (in main goroutine)
//		for x:=range squares{
//			fmt.Println(x)
//		}
//}

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	//8.4.3
	//naturals :=make(chan int)
	//squarers :=make(chan int)
	//go counter(naturals)
	//go squarer(squarers,naturals)
	//printer(squarers)

	//8.4.4
	ch := make(chan string, 3)
	ch <- "A"
	ch <- "B"
	ch <- "C"
	fmt.Println(<-ch)    // "A"
	fmt.Println(cap(ch)) // "3"
	fmt.Println(len(ch)) // "2"
	fmt.Println(<-ch)    // "B"
	fmt.Println(<-ch)    // "C"
}
