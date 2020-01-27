/*
 *Author: Stefan
 *Date: 01/27/2020
 *Last changes: 01/27/2020 02:22
 *Task: Class Work L3 Ex 1-15
 */

package main

import "fmt"

//main func
func main() {
	//ex 1
	/*
		var x [5]int
		fmt.Println(x)
		x[1] = 4
		fmt.Println(x)
	*/
	//ex 2
	/*
		var a = [...]int{1, 2, 3, 4}
		fmt.Println(a, len(a))
	*/
	//ex 3
	/*
		arr := [5]int{0: 1, 2, 3, 4}
		fmt.Println(arr, len(arr))
		arr1 := [5]int{0: 1, 1: 1, 2: 1, 3: 1}
		fmt.Println(arr1, len(arr1))
	*/
	//ex 4
	/*
		a := [5]int{2: 3, 0: 1, 2}
		fmt.Println(a, len(a))
	*/
	//ex 5
	/*
		var x [5]int
		fmt.Println(x[0], x[4])
	*/
	//ex 6
	/*
		var (
			a, y [5]int
			x    [6]int
		)
		fmt.Println(a == y)
		fmt.Printf("a, y: %v %T | x: %v %T \n\r\n", a, a, x, x)
	*/
	//ex 7
	/*
		arr := [3]string{
			"cat",
			"dog",
			"fox",
		}
		for i, v := range arr {
			fmt.Println(i, v)
		}
		for i := 0; i < len(arr); i++ {
			fmt.Println(i, arr[i])
		}
	*/
	//ex 8
	/*
		a := [3]int{1, 2, 3}
		fmt.Println(a)
		chang(a)
		fmt.Println(a)
	*/
	//ex 9
	/*
		a := [2][2]int{
			[2]int{2, 1},
			[2]int{1, 2},
		}
		fmt.Println(a)
		fmt.Println(a[0][0], a[0][1])
		fmt.Println(a[1][0], a[1][1])
	*/
	//ex 10
	/*
		var a []int
		var b = []int{1, 2, 3}

		c := []bool{true, false}
		d := []string{
			"one",
			"four",
		}

		e := make([]int, 5) // return slice
		f := make([]int, 3, 7)
		g := new([]int) //return pointer

		fmt.Printf("a: %v %T| len: %v | cap: %v \n\r\n", a, a, len(a), cap(a))
		fmt.Printf("b: %v %T| len: %v | cap: %v \n\r\n", b, b, len(b), cap(b))
		fmt.Printf("c: %v %T| len: %v | cap: %v \n\r\n", c, c, len(c), cap(c))
		fmt.Printf("d: %v %T| len: %v | cap: %v \n\r\n", d, d, len(d), cap(d))
		fmt.Printf("e: %v %T| len: %v | cap: %v \n\r\n", f, f, len(f), cap(f))
		fmt.Printf("f: %v %T| len: %v | cap: %v \n\r\n", e, e, len(e), cap(e))
		fmt.Printf("g: %v %T| \n\r\n", g, g)
	*/
	//ex 11
	/*
		s := make([]int, 6)
		fmt.Println(s, len(s), cap(s))
		s = s[2:6] // 2 3 4 5 indexes
		fmt.Println(s, len(s), cap(s))
	*/
	//ex 12
	/*
		s := make([]int, 6)
		fmt.Println(s, len(s), cap(s)) //0 1 2 3 4 5
		//slicing
		s = s[2:4] // 2 3 but cap devide one target array
		fmt.Println(s, len(s), cap(s))
		s = s[:cap(s)]
		fmt.Println(s, len(s), cap(s))
	*/
	//ex 13
	/*
		m := []int{1, 2, 3, 4, 5, 6}
		fmt.Println(m, len(m), cap(m))
		q1 := m[0:3]
		fmt.Println(q1, len(q1), cap(q1))
		q2 := m[2:4]
		fmt.Println(q2, len(q2), cap(q2))
		m[2] = 55

		fmt.Println()
		fmt.Println(m, len(m), cap(m))
		fmt.Println(q1, len(q1), cap(q1))
		fmt.Println(q2, len(q2), cap(q2))
		q2 = append(q2, 1, 2, 3, 4) // this slice are grown ang now it don;t
		//linking on m, now it is a adult slice with self target array

		fmt.Println()
		fmt.Println(m, len(m), cap(m))
		fmt.Println(q1, len(q1), cap(q1))
		fmt.Println(q2, len(q2), cap(q2))

		m[2] = 666

		fmt.Println()
		fmt.Println(m, len(m), cap(m))
		fmt.Println(q1, len(q1), cap(q1))
		fmt.Println(q2, len(q2), cap(q2))
	*/
	//ex 14 Growing slices experiment
	/*
		sl := make([]int, 1)
		for i := 1; i < 1026; i++ {
			fmt.Printf("len: %3v | cap: %3v | sl: %v\n\r", len(sl), cap(sl), sl)
			sl = append(sl, i)
		}
	*/
	//ex 15
	/*
		//slices are reference type
		s := []int{1, 2, 3}
		fmt.Println(s)
		changeSlice(s)
		fmt.Println(s)
		fmt.Println(s)
	*/
	fmt.Println()

	//ex 16

}

/*
//ex 8 //work only via values
//chang func
func chang(a [3]int) {
	a[1] = 777
	fmt.Println(a)
}
*/

//ex 15
/*
//changeSlice func
func changeSlice(s []int) {
	if len(s) > 0 {
		s[0] = 88
	}
}
*/
