/*
 *Author: Stefan
 *Date: 02/05/2020
 *Last changes: 02/06/2020 15:25
 *Task: Concurrency 4 Patterns
 */

package main

import (
	"fmt"
	"math/rand"
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

	signalClose()
	fmt.Println()

	signalAck()
	fmt.Println()

	closeRange()
	fmt.Println()

	selectRecv()
	fmt.Println()

	selectSend()
	fmt.Println()

	selectDrop()
	fmt.Println()
	fmt.Println()

	//ex 4 LOOP4
	/*
	   	//create chan buff size 2
	   	ch1 := make(chan int, 2)
	   	//send val
	   	ch1 <- 1
	   	ch1 <- 2

	   	ch2 := make(chan int, 2)
	   	ch2 <- 3
	   	//for loop with ability to stop it
	   	//scheduler make decission who first and who sec
	   LOOP:
	   	for {
	   		select {
	   		case v1 := <-ch1:
	   			fmt.Println(v1, "chan1 val")
	   		case v2 := <-ch2:
	   			fmt.Println(v2, "chan2 val")
	   		default:
	   			break LOOP
	   		}
	   	}
	*/
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

//
func signalClose() {
	ch := make(chan struct{})
	go func() {
		time.Sleep(100 * time.Microsecond)
		fmt.Println("signal event")
		close(ch)
	}()
	<-ch
	fmt.Println("event received")
}

//signalAck func
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

//closeRange func
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

//selectSend func
func selectSend() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	}()

	select {
	case ch <- "work":
		fmt.Println("send work")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("time out")
	}
}

//selectDrop func
func selectDrop() {
	ch := make(chan int, 5)

	go func() {
		for v := range ch {
			fmt.Println("recv", v)
		}
	}()

	for i := 0; i < 20; i++ {
		select {
		case ch <- i:
		default:
			fmt.Println("drop", i)
		}
	}
	close(ch)
	fmt.Scanln()
}
