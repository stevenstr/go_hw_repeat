/*
 *Author: Stefan
 *Date: 03/10/2020
 *Last changes: 03/10/2020 22:53
 *Task: Chapter 3 Lecture 3 HW 3.1 AVG
 */

package main

import (
	"fmt"

	"github.com/stevenstr/go_hw_repeat/avg"
)

//main func
func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(avg.Average(arr))
}
