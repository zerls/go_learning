package main

import (
	"fmt"
	"time"
)

//1. Channels
// func main() {
// 	messages := make(chan string)

// 	go func() {
// 		messages <- "ping"
// 	}()

// 	msg := <-messages
// 	fmt.Println(msg)
// }

// 2. Channel Buffering
// func main() {
// 	messages := make(chan string, 2)

// 	messages <- "buffered"
// 	messages <- "channel"

// 	fmt.Println(<-messages)
// 	fmt.Println(<-messages)
// }

// 3. Channel Synchronization
// func worker(done chan bool) {
// 	fmt.Print("Working...")
// 	time.Sleep(time.Second)
// 	fmt.Print("done")

// 	done <- true
// }

// func main() {
// 	done := make(chan bool, 1)
// 	go worker(done)

// 	<-done
// }

// 4. Channel Directions
// func ping(pings chan<- string, msg string) {
// 	pings <- msg
// }
// func pong(pings <-chan string, pongs chan<- string) {
// 	msg := <-pings
// 	pongs <- msg
// }

// func main() {
// 	pings := make(chan string, 1)
// 	pongs := make(chan string, 1)
// 	ping(pings, "passed message")
// 	pong(pings, pongs)
// 	fmt.Println(<-pongs)
// }

// 5. Select
// func main() {
// 	c1 := make(chan string)
// 	c2 := make(chan string)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		c1 <- "one"
// 	}()

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		c2 <- "two"
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case msg1 := <-c1:
// 			fmt.Println("received", msg1)
// 		case msg2 := <-c2:
// 			fmt.Println("received", msg2)
// 		}
// 	}
// }

//6. Timeouts

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

// // 6. Timer
// func main() {
// 	time1 := time.NewTimer(2 * time.Second)

// 	<-time1.C
// 	fmt.Println("Time 1 fired")

// 	time2 := time.NewTimer(time.Second)

// 	go func() {
// 		<-time2.C
// 		fmt.Println("Time 2 fired")
// 	}()
// 	stop2 := time2.Stop()
// 	// stop2 := false
// 	if stop2 {
// 		fmt.Println("Time 2 stoped")
// 	}

// 	time.Sleep(2 * time.Second)
// }

// 7. Ticker

// func main() {
// 	ticker := time.NewTicker(500 * time.Millisecond)
// 	done := make(chan bool)

// 	go func() {
// 		for {
// 			select {
// 			case <-done:
// 				return
// 			case t := <-ticker.C:
// 				fmt.Println("tick at", t)
// 			}
// 		}
// 	}()

// 	time.Sleep(1600 * time.Millisecond)
// 	ticker.Stop()
// 	done <- true
// 	fmt.Println("Ticker stoped")
// }
