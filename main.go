/*
 *Author: Stefan
 *Date: 01/27/2020
 *Last changes: 01/27/2020 22:13
 *Task: HW 3.1 Average
 */

package main

import "fmt"

//average func
func average(arr [6]int) float64 {
	sum := 0
	for _, v := range arr {
		sum = sum + v
	}
	return float64(sum) / float64(len(arr))
}

//main func
func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(average(arr))
}
