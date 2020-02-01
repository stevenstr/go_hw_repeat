/*
 *Author: Stefan
 *Date: 02/01/2020
 *Last changes: 02/01/2020 19:35
 *Task: HW 3.4 Maps
 */

package main

import (
	"fmt"
	"sort"
)

//printSorted func
func printSorted(m map[int]string) {
	sl := []int{}
	for k := range m {
		sl = append(sl, k)
	}
	sort.Ints(sl)
	for _, i := range sl {
		fmt.Printf(m[i] + " ")
	}
	fmt.Println()
}

//main func
func main() {
	m := map[int]string{1: "1", 7: "7", 5: "5"}
	fmt.Println(m)
	printSorted(m)

	m1 := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	fmt.Println(m1)
	printSorted(m1)

	m2 := map[int]string{2: "a", 0: "b", 1: "c"}
	fmt.Println(m2)
	printSorted(m2)
}
