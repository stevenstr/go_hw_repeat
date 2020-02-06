/*
 *Author: Stefan
 *Date: 02/06/2020
 *Last changes: 02/06/2020 16:27
 *Task: Concurrency 7 Mutex
 */

package main

import (
	"fmt"
	"sync"
)

//main finc
func main() {
	var counters = map[int]int{}
	mu := &sync.Mutex{}
	for i := 0; i < 5; i++ {
		go func(counters map[int]int, th int, mu *sync.Mutex) {
			for j := 0; j < 5; j++ {
				//fix race condition
				mu.Lock()
				counters[th*10+j]++
				mu.Unlock()
			}
		}(counters, i, mu)
	}
	//wait some time
	//fmt.Scanln()

	//fix race condition
	mu.Lock()
	fmt.Println(counters)
	mu.Unlock()
}
