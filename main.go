/*
 *Author: Stefan
 *Date: 02/05/2020
 *Last changes: 02/06/2020 15:19
 *Task: Concurrency 4
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	//ex 1
	/*
		//unbuffereed
		//ch1:= make(chan int)
		//                   , 0
		//buffered channel
		ch := make(chan int, 2)
		//send to chan 2
		ch <- 2
		ch <- 44
		// deadlock on send
		//ch <- 5
		//read from channel buffer
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		//deadlock in read
		//fmt.Println(<-ch)
	*/
	//ex 2
	/*
		//unbuffered
		//ch := make(chan int, 0)
		//buffered
		ch := make(chan int, 1)
		//
		go func(c chan int) {
			fmt.Println("before reading")
			fmt.Println(<-c)
			fmt.Println("after reading")
		}(ch)
		fmt.Println("before sending")
		ch <- 2
		fmt.Println("after sending")
		fmt.Scanln() // lock
	*/

	//patterns
	//ex 3
	basicSR()
	fmt.Println()

	sigClose()
	fmt.Println()

	signalAck()
	fmt.Println()

	closeRange()
	fmt.Println()

	selectRecv()

}

func basicSR() {
	ch := make(chan string)
	//child send
	go func() {
		ch <- "pay"
		close(ch)
	}()
	//main read
	fmt.Println(<-ch)
}

func sigClose() {
	ch := make(chan struct{})
	go func() {
		time.Sleep(100 * time.Microsecond)
		fmt.Println("signal event")
		close(ch)
	}()
	<-ch
	fmt.Println("event received")
}

func signalAck() {
	ch := make(chan string)
	//child gorout
	go func() {
		fmt.Println(<-ch)
		ch <- "ok done"
	}()
	//main goroutine
	ch <- "do this"
	fmt.Println(<-ch)
}

func closeRange() {
	//create box with 5 int
	ch := make(chan int, 5)
	//send data (5 ints)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	//close channel
	close(ch)

	//if channel close - read all from channel and after stop it
	/*for v := range ch {
		fmt.Println(v)
	}*/
}

//select multyplexor
func selectRecv() {
	ch := make(chan int)
	ch2 := make(chan int, 1)

	select {
	case v := <-ch:
		fmt.Println(v, "ch value")
	case ch2 <- 1:
		fmt.Println("put value in ch2")
	default:
		fmt.Println("Default case")
	}
}
