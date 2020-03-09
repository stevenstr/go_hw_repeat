/*
 *Author: Stefan
 *Date: 03/09/2020
 *Last changes: 03/09/2020 13:06
 *Task: Chapter 1 Lecture 1 HW 1.1
 */

package main

import "fmt"

//main func
func main() {
	sl := []string{"H", "e", "l", "l", "o", " ", "w", "o", "r", "l", "d", "!"}

	for i := 0; i < len(sl); i++ {
		fmt.Printf(sl[i])
	}
	fmt.Println()

	for _, v := range sl {
		fmt.Printf(v)
	}
	fmt.Println()
}
