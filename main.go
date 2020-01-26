/*
 *Author: Stefan
 *Date: 01/27/2020
 *Last changes: 01/27/2020 01:-2
 *Task: HW 2.1 Factorial
 */

package main

import "fmt"

//factorial func
func factorial(i uint) uint {
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
	fmt.Println(factorial(7))
	fmt.Println(factorial(1))
	fmt.Println(factorial(0))
}
