/*
 *Author: Stefan
 *Date: 02/05/2020
 *Last changes: 02/06/2020 13:56
 *Task: Concurrency 3 Checker Self made
 */

package main

import (
	"fmt"
	"net/http"
	"time"
)

//main func
func main() {

	//sources
	res := []string{
		"http://epam.com",
		"http://tut.by",
		"http://google.com",
	}

	//channel to communicate
	c := make(chan string)

	//start goroutines
	for _, url := range res {
		go checkRes(url, c)
	}

	//read from channel
	for l := range c {
		go func(url string) {
			time.Sleep(5 * time.Second)
			checkRes(url, c)
		}(l)
	}
}

//checkRes func
func checkRes(url string, c chan string) {
	//Get
	_, err := http.Get(url)
	//error handler
	if err != nil {
		fmt.Println(url, "down")
		c <- url
		return
	}
	fmt.Println(url, "up")
	c <- url
	return
}
