/*
 *Author: Stefan
 *Date: 03/09/2020
 *Last changes: 03/09/2020 17:42
 *Task: Chapter 2 Lecture 2 HW 2.1 factorial
 */

package main

import "fmt"

func main() {
	fmt.Println(factorial(0))
	fmt.Println(factorial(1))
	fmt.Println(factorial(3))
	fmt.Println(factorial(4))
	fmt.Println(factorial(5))
	fmt.Println(factorial(6))
}

//factorial func
func factorial(a uint) (res uint) {
	if a == 0 {
		return 1
	}
	res = 1
	for i := uint(2); i <= a; i++ {
		res = res * i
	}
	return
}
