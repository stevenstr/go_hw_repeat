/*
 *Author: Stefan
 *Date: 01/26/2020
 *Last changes: 01/26/2020 23:35
 *Task: HW 1.1 Hello World
 */

package main

import (
	"fmt"
	"unicode/utf8"
)

func max(sl []string) string {
	res := ""

	for i := 0; i < len(sl); i++ {
		if utf8.RuneCountInString(sl[i]) > utf8.RuneCountInString(res) {
			res = sl[i]
		}
	}
	return res
}

func reverse(sl []int64) []int64 {
	rsl := []int64{}
	for i := len(sl) - 1; i >= 0; i-- {
		rsl = append(rsl, sl[i])
	}
	return rsl
}

func main() {
	sl := []string{"three", "one", "twelve", "stefan"}
	sl1 := []int64{21, 41, 5, 1}

	fmt.Println(max(sl))
	fmt.Println(reverse(sl1))
}
