/*
 *Author: Stefan
 *Date: 02/06/2020
 *Last changes: 02/06/2020 17:35
 *Task: HW 5.2 Testing (unit testing)
 */

package main

import (
	"fmt"

	c "github.com/stevenstr/go_hw_repeat/conv"
)

func main() {
	sl := []string{"11111", "dffs", "1g1g1", ""}

	for _, v := range sl {
		r, err := c.MyStrToIntAtoi(v)
		fmt.Println(r, err)
	}
}
