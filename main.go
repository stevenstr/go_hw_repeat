/*
 *Author: Stefan
 *Date: 01/26/2020
 *Last changes: 01/26/2020 23:35
 *Task: HW 1.1 Hello World
 */

package main

import "fmt"

//factorial func
func factorial(i uint) uint {
	if i < 0 {
		fmt.Println("Negative number!")
		return 0
	}
	var res, a uint
	res, a = 1, 1
	for a = 1; a <= i; a++ {
		res *= a
	}
	return res
}

//main func
func main() {
	fmt.Println(factorial(5))
}
