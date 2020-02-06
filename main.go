/*
 *Author: Stefan
 *Date: 02/06/2020
 *Last changes: 02/06/2020 16:33
 *Task: Concurrency 8 Atomic
 */

package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

//totalOperations var
var totalOperations int32

//inc func
func inc() {
	atomic.AddInt32(&totalOperations, 1)
}

//main func
func main() {
	for i := 0; i < 1000; i++ {
		go inc()
	}
	time.Sleep(2 * time.Millisecond)
	fmt.Println("Score: ", totalOperations)
}
