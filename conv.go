/*
 *Author: Stefan
 *Date: 02/05/2020
 *Last changes: 02/05/2020 22:35
 *Task: HW 5.1 Testing (unit testing)
 */

package main

import (
	"fmt"
	"strconv"
)

//MyStrToInt func
func MyStrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func main() {
	sl := []string{"11111", "dffs", "1g1g1", ""}

	for _, v := range sl {
		r, err := MyStrToInt(v)
		fmt.Println(r, err)
	}
}
