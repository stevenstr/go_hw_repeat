/*
 *Author: Stefan
 *Date: 02/06/2020
 *Last changes: 02/06/2020 16:12
 *Task: Concurrency 6 Wait Group
 */

package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

//concst
const (
	iter         = 7
	goroutNumber = 5
)

//startWorker
func startWorker(in int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < iter; j++ {
		fmt.Printf(formatWork(in, j))
	}
}

//main func
func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < goroutNumber; i++ {
		wg.Add(1)
		go startWorker(i, wg)
	}
	time.Sleep(time.Millisecond)

	wg.Wait()
}

//formatWork func
func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat(" ", in), "#",
		strings.Repeat(" ", goroutNumber-in),
		"th", in,
		"iter", j, strings.Repeat("*", j))
}
