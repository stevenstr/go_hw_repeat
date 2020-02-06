/*
 *Author: Stefan
 *Date: 02/06/2020
 *Last changes: 02/06/2020 13:35
 *Task: example goroutines
 */

package main

import (
	"fmt"
	"net/http"
	"time"
)

//main func
func main() {
	resources := []string{
		"http://google.by",
		"http://tut.by",
		"http://vk.com",
	}

	//create channel for communicate between goroutines
	c := make(chan string)

	//just run child goroutines
	for _, url := range resources {
		go checkRes(url, c)
	}

	//read
	for l := range c {
		go func(url string) {
			time.Sleep(5 * time.Second)
			checkRes(url, c)
		}(l)
	}
}

//checkRes func
func checkRes(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down!")
		c <- url
		return
	}

	fmt.Println(url, "is Up!")
	c <- url
	return
}
