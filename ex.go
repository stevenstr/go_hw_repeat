/*
 *Author: Stefan
 *Date: 02/05/2020
 *Last changes: 02/05/2020 22:35
 *Task: Class Work L5 Testing (unit testing, )
 */

package main

import "fmt"

func summ(a []int) (result int) {
	for _, v := range a {
		result += v
	}
	return result
}

func dummy(k int) (int, error) {
	if k > 10 {
		return 1, fmt.Errorf("more")
	} else if k < 10 {
		return -1, fmt.Errorf("less")
	}
	return 0, fmt.Errorf("zero")
}

func main() {
}
