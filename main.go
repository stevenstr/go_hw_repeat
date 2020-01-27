/*
 *Author: Stefan
 *Date: 01/28/2020
 *Last changes: 01/28/2020 02:32
 *Task: HW 3.2 Max
 */

package main

import (
	"fmt"
	"unicode/utf8"
)

//Max func
func Max(st []string) string {
	l := len(st) // 4
	res := ""
	for i := 0; i < l; i++ {
		if utf8.RuneCountInString(st[i]) > utf8.RuneCountInString(res) {
			res = st[i]
		}
	}
	return res
}

//main func
func main() {
	sl := []string{"a", "bbb", "cc", "lesha", "ana"}
	fmt.Println(Max(sl))
	sl1 := []string{"ssss", "bbb", "cgggggg", "lesha"}
	fmt.Println(Max(sl1))
}
