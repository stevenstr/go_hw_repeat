/*
 *Author: Stefan
 *Date: 01/28/2020
 *Last changes: 01/28/2020 02:42
 *Task: HW 3.3 Reverse
 */

package main

import "fmt"

//Reverse func
func Reverse(a []int64) []int64 {
	res := make([]int64, 0)
	for i := len(a) - 1; i >= 0; i-- {
		res = append(res, a[i])
	}
	return res
}

//main func
func main() {
	sl := []int64{1, 2, 4, 66}
	sl1 := []int64{2, 4, 6, 9}
	sl2 := []int64{55, 22, 54, 7}

	fmt.Println(Reverse(sl))
	fmt.Println(Reverse(sl1))
	fmt.Println(Reverse(sl2))
}
