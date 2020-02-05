/*
 *Author: Stefan
 *Date: 02/05/2020
 *Last changes: 02/05/2020 20:59
 *Task: HW 4.1 Age Name
 */

package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

//Person struct
type Person struct {
	firstName string
	lastName  string
	birthDay  time.Time
}

//People slice of struct
type People []Person

//Len method
func (p People) Len() int {
	return len(p)
}

//Swap method
func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//Less method
func (p People) Less(i, j int) bool {

	if p[i].birthDay.Sub(p[j].birthDay) > 0 {
		return true
	}
	if p[i].birthDay.Sub(p[j].birthDay) < 0 {
		return false
	}
	if p[i].firstName < p[j].firstName {
		return true
	}
	if p[i].firstName > p[j].firstName {
		return false
	}
	return p[i].lastName < p[j].lastName
}

//main func
func main() {
	kolyaIvanovDate, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ivanIvanovDate, err := time.Parse("2006-Jan-02", "2003-Aug-05")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	LeshaIvanovDate, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//create array of test data
	p := People{
		{"Kolya", "Ivanov", kolyaIvanovDate},
		{"Ivan", "Ivanov", ivanIvanovDate},
		{"Lesha", "Ivanov", LeshaIvanovDate},
	}

	//sorting
	sort.Sort(p)

	//show results
	for _, q := range p {
		fmt.Println(q.firstName, q.lastName, q.birthDay)
	}
}
