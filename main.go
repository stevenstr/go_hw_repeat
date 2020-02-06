/*
 *Author: Stefan
 *Date: 02/05/2020
 *Last changes: 02/06/2020 12:57
 *Task: Concurrency 1 Checker
 */

package main

import (
	"fmt"
	"net/http"
)

func main() {
	// main rourine
	resources := []string{
		"http://tut.by",
		"http://google.by",
		"http://golang.com",
		"http://onliner.by",
		"http://github.com",
	}
	//create chanal for communicate using string
	c := make(chan string)

	//cicle start in main goroutine
	for _, url := range resources {
		go checkRes(url, c) //create new goroutine
	}
	//waiting value from channel
	/*
		fmt.Println(<-c) //this operation are blocking!! and print facter answer
		fmt.Println(<-c)
		fmt.Println(<-c)
		fmt.Println(<-c)
		fmt.Println(<-c)
		//fmt.Println(<-c)// dead lock of main routine
	*/

	//waiting value from channel
	for range resources {
		fmt.Println(<-c)
	}
}

func checkRes(url string, c chan string) {
	_, err := http.Get(url) //processor wait, GET is block operation
	if err != nil {
		//send to channel
		c <- fmt.Sprintln(url, "is down!")
		return
	}
	//send to channel
	c <- fmt.Sprintln(url, "is Up!")
	return
}
