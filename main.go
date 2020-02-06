/*
 *Author: Stefan
 *Date: 02/05/2020
 *Last changes: 02/06/2020 15:35
 *Task: Concurrency 5
 */

package main

import "fmt"

//main func
func main() {
	//chan for undo
	cancelCh := make(chan struct{})
	//chan for data
	dataCh := make(chan int)

	go func(cancelCh chan struct{}, dataCh chan int) {
		val := 0
		for {
			select {
			case <-cancelCh:
				close(dataCh)
				return
			case dataCh <- val:
				val++
			}
		}
	}(cancelCh, dataCh)

	//read data channel from child goroutine
	for curVal := range dataCh {
		fmt.Println("read", curVal)
		if curVal > 3 {
			fmt.Println("send cancel")
			//send cancelch for case in select
			cancelCh <- struct{}{}
		}
	}
}
