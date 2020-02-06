/*
 *Author: Stefan
 *Date: 02/05/2020
 *Last changes: 02/06/2020 13:23
 *Task: Concurrency 2 Checker
 */

package main

import (
	"fmt"
	"net/http"
	"time"
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
	//just init
	for _, url := range resources {
		go checkRes(url, c) //create new goroutine
	}
	//read from channel
	//and create new goroutines^))))))
	/*for {
		go checkRes(<-c, c)
	}
	*/
	//alternative way
	/*
		for l := range c {
			checkRes(l, c)
		}
	*/
	for l := range c {
		//bad method, we just stopped main routine
		//time.Sleep(5 * time.Second)
		//send parameter so that linking woll be on diff mamory place
		go func(url string) {
			time.Sleep(5 * time.Second)
			checkRes(url, c)
		}(l)
	}
}

//checkRes func
func checkRes(url string, c chan string) {

	_, err := http.Get(url) //processor wait, GET is block operation
	if err != nil {
		//send to channel
		fmt.Println(url, "is down!")
		c <- url
		return
	}
	//send to channel
	fmt.Println(url, "is Up!")
	c <- url
	return
}
