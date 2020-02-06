/*
 *Author: Stefan
 *Date: 02/05/2020
 *Last changes: 02/05/2020 22:35
 *Task: HW 5.1 Testing (unit testing)
 */

package main

import (
	"fmt"

	c "github.com/stevenstr/go_hw_repeat/conv"
)

func main() {
	sl := []string{"11111", "dffs", "1g1g1", ""}

	for _, v := range sl {
		r, err := c.MyStrToInt(v)
		fmt.Println(r, err)
	}
}
