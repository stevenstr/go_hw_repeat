/*
 *Author: Stefan
 *Date: 03/10/2020
 *Last changes: 03/10/2020 22:27
 *Task: Chapter 3 Lecture 3 Arrays, Slices, Maps ex 1-31
 */

package main

//import "fmt"

//main func
func main() {

	//ex 1
	/*
		arr := [6]int{}
		fmt.Println(arr)
		fmt.Println()

		for i, v := range arr {
			fmt.Printf("index: %v | value: %v | type: %T \n\r\n", i, v, v)
		}
		fmt.Println()

		arr1 := [3]int{1, 2, 3}
		fmt.Println(arr1)
		fmt.Println()

		for i, v := range arr1 {
			fmt.Printf("index: %v | value: %v | type: %T \n\r\n", i, v, v)
		}
	*/

	//ex 2
	/*
		var arr [6]int

		var arr1 = [4]float32{}

		var arr2 = [4]string{"a", "b", "c", "d"}

		var arr3 = [2][2]int{
			[2]int{1, 2},
			[2]int{3, 4},
		}

		fmt.Println(arr)
		fmt.Println(arr1)
		fmt.Println(arr2)
		fmt.Println(arr3)

		arr4 := [3]int{}
		arr5 := [3]int{
			1,
			2,
			3,
		}

		fmt.Println(arr4)
		fmt.Println(arr5)
	*/

	//ex 3
	/*
		var x [3]int
		fmt.Println(x)

		var y = [3]int{}
		fmt.Println(y)

		y[1] = 2
		y[0] = 1
		y[2] = 3

		fmt.Println(y)
	*/

	//ex 4 auto length declaration
	/*
		foo := [...]int{1, 2, 3, 4, 5, 6, 7}
		bar := [...]int{
			1,
			2,
			3,
		}

		fmt.Println(foo, bar)
	*/

	//ex 5
	/*
		foo := [...]int{1, 2, 3, 4}

		fmt.Println(foo, "len: ", len(foo))
	*/

	//ex 6
	/*
		foo := [3]string{"hell", "my", "eye"}
		fmt.Println(foo)
		fmt.Println("len: ", len(foo))
		fmt.Printf("%T\n\r\n", foo)
	*/

	//ex 7 array literals
	/*
		foo := [6]int{2: 3, 4, 5, 6} //0 1 2 3 4 5
		fmt.Println(foo)

		foo[0] = 1
		foo[1] = 2
		fmt.Println(foo)
	*/

	//ex 8 array literals
	/*
		foo := [6]int{3: 4, 5, 6, 0: 1, 2, 3}

		fmt.Println(foo)
		fmt.Println("len: ", foo)

		fmt.Printf("%T\n\r\n", foo)
	*/

	//ex 9
	/*
		x, y := [4]int{}, [4]int{}
		fmt.Println(x, y)

		fmt.Printf("x: %T | y: %T\n\r\n", x, y)
		fmt.Println("x == y? ", x == y)

		x1, y1 := [5]int{}, [4]int{}
		fmt.Println(x, y)

		fmt.Printf("x1: %T | y1: %T\n\r\n", x, y)
		//fmt.Println("x1 == y1? ", x1 == y1) //mismatched types
	*/

	//ex 10 array are not linked type
	/*
		arr := [4]int{}
		fmt.Println(arr)

		foo(arr)

		fmt.Println(arr)
	*/

	//ex 11 arrray iteration using len or range
	/*
		var arr [10]int

		fmt.Println("Using len: ")
		//using len
		for i := 0; i < len(arr); i++ {
			fmt.Println("index: ", i, " | value:", arr[i])
		}
		fmt.Println()

		fmt.Println("using range: ")
		//using range
		for i, v := range arr {
			fmt.Println("index: ", i, " | value:", v)
		}
	*/

	//ex 12
	/*
		arr := [10]int{}

		for _, v := range arr {
			fmt.Println("value: ", v)
		}

		arr1 := [10]int{}

		for i := range arr1 {
			fmt.Println("index: ", i)
		}
	*/

	//SLICES
	//ex 13
	/*
		var sl []int
		var sl1 = []int{}
		var sl2 = []int{
			1,
			2,
			3,
		}

		fmt.Printf("%v | %T\n\r\n", sl, sl)
		fmt.Printf("%v | %T\n\r\n", sl1, sl1)
		fmt.Printf("%v | %T\n\r\n", sl2, sl2)
	*/

	//ex 14 make new
	/*
		sl := make([]int, 5)
		sl1 := make([]int, 3, 10)

		fmt.Println("value: ", sl, "| len: ", len(sl), "| cap: ", cap(sl))
		fmt.Println("value: ", sl1, "| len: ", len(sl1), "| cap: ", cap(sl1))

		sl2 := new([]int)
		fmt.Println("value: ", sl2)
	*/

	//ex 15
	/*
		var a []int
		var b = []int{1, 2, 3}

		c := []bool{true, true}
		d := []string{"hello, ", "world", "!"}

		e := make([]int, 5)
		f := make([]int, 5, 10)

		g := new([]string)

		fmt.Println(a, b, c, d, e, f, g)

		fmt.Printf("Type: %T | Value: %v\n\r\n", a, a)
		fmt.Printf("Type: %T | Value: %v\n\r\n", b, b)
		fmt.Printf("Type: %T | Value: %v\n\r\n", c, c)
		fmt.Printf("Type: %T | Value: %v\n\r\n", d, d)
		fmt.Printf("Type: %T | Value: %v\n\r\n", e, e)
		fmt.Printf("Type: %T | Value: %v\n\r\n", g, g)
	*/

	//ex 16 Slicing
	/*
		sl := make([]int, 6) //len: 6, cap: 6
		fmt.Printf("value: %v | len: %v | cap: %v \n\r\n", sl, len(sl), cap(sl))

		for i := 0; i < len(sl); i++ {
			sl[i] = (i + 1)
		}

		fmt.Printf("value: %v | len: %v | cap: %v \n\r\n", sl, len(sl), cap(sl))

		sl = sl[1:4] //1, 2, 3 len: 3, cap: 5
		fmt.Printf("value: %v | len: %v | cap: %v \n\r\n", sl, len(sl), cap(sl))

		//Расщиряем cap
		sl = sl[:cap(sl)] // len: 5, cap: 5
		fmt.Printf("value: %v | len: %v | cap: %v \n\r\n", sl, len(sl), cap(sl))
	*/

	//ex 17 Slicing example
	/*
		sl := []int{1, 2, 3, 4, 5, 6}
		fmt.Printf("value: %v | len: %v | cap: %v \n\r", sl, len(sl), cap(sl))
		q := sl[1:5]
		fmt.Printf("value: %v | len: %v | cap: %v \n\r", q, len(q), cap(q))
		w := sl[3:]
		fmt.Printf("value: %v | len: %v | cap: %v \n\r\n", w, len(w), cap(w))

		fmt.Println()
		//affect on all
		q[3] = 1000000

		fmt.Printf("value: %v | len: %v | cap: %v \n\r", sl, len(sl), cap(sl))
		fmt.Printf("value: %v | len: %v | cap: %v \n\r", q, len(q), cap(q))
		fmt.Printf("value: %v | len: %v | cap: %v \n\r\n", w, len(w), cap(w))

		//no affect because cap overwrite
		q = append(q, sl...)
		fmt.Printf("value: %v | len: %v | cap: %v \n\r", sl, len(sl), cap(sl))
		fmt.Printf("value: %v | len: %v | cap: %v \n\r", q, len(q), cap(q))
		fmt.Printf("value: %v | len: %v | cap: %v \n\r\n", w, len(w), cap(w))
	*/

	//ex 18
	/*
		sl := make([]int, 1, 5)
		fmt.Println(sl, len(sl), cap(sl))

		sl = append(sl, 1, 2, 3, 4)
		fmt.Println(sl, len(sl), cap(sl))
	*/

	//ex 19 aafection with growing slice
	/*
		arr := [...]int{1, 2, 3, 4}
		sl := arr[1:3]
		fmt.Println("array len: ", len(arr))
		fmt.Println("array: ", arr, "sl:", sl, "len:", len(sl), "cap:", cap(sl))
		fmt.Println()

		arr[2] = 333
		fmt.Println("array: ", arr, "sl:", sl, "len:", len(sl), "cap:", cap(sl))
		sl = append(sl, 66)
		fmt.Println("array: ", arr, "sl:", sl, "len:", len(sl), "cap:", cap(sl))
		sl = append(sl, 77)
		fmt.Println("array: ", arr, "sl:", sl, "len:", len(sl), "cap:", cap(sl))
		sl = append(sl, 88)
		fmt.Println("array: ", arr, "sl:", sl, "len:", len(sl), "cap:", cap(sl))
	*/

	//ex 20 GROWING CAPACITY
	/*
		sl := make([]int, 1)

		for i := 1; i < 68; i++ {
			sl = append(sl, 1)

			fmt.Println("len:", len(sl), "cap:", cap(sl), sl)

			if len(sl) == cap(sl) {
				fmt.Println()
				fmt.Println("GROWING")
				fmt.Println()
			}
		}
	*/

	//ex 21 Slice are refs type
	/*
		sl := []int{1, 2, 3}
		fmt.Println(sl)
		ref(sl)
		fmt.Println(sl)
	*/

	//MAPS
	//ex 22
	/*
		var m map[int]int
		var m1 = map[int]int{1: 1, 2: 2, 3: 3}

		fmt.Println(m, m1)

		m2 := map[int]bool{
			1: true,
			2: true,
		}
		m3 := map[bool]string{true: "yes", false: "no!"}

		fmt.Println(m2, m3)
	*/

	//ex 23
	/*
		m := make(map[bool]int, 2)
		fmt.Println(m)

		m1 := make(map[int]int, 2)
		fmt.Println(m1)
	*/

	//ex 24
	/*
		m := map[int]string{}

		m[1] = "one"
		m[3] = "three"

		fmt.Println(m)

		m[3] = "Reva!"
		fmt.Println(m)
	*/

	//ex 25
	/*
		m := map[string]int{}

		m["one"] = 1
		m["two"] = 2

		a := m["one"]
		b := m["two"]

		//if no exist return default value for type
		c := m["three"]

		fmt.Println(m)
		fmt.Println(a, b, c)
	*/

	//ex 26 check elements
	/*
		m := map[int]string{1: "one", 2: "two"}

		one, ok := m[1]
		if !ok {
			fmt.Println(one, "is not exist!")
		}

		two, ok := m[2]
		if !ok {
			fmt.Println(two, " is not exist!")
		}

		three, ok := m[3]
		if !ok {
			fmt.Println(three, "three is not exist!")
		}
	*/

	//ex 27 nil map
	/*
		var m map[int]int

		a, ok := m[1]
		if !ok {
			fmt.Println(a, ok)
		}
	*/

	//ex 28 delete elements
	/*
		m := map[string]int{"one": 1, "two": 2, "zero": 3}

		fmt.Println(m, len(m))

		delete(m, "zero")

		fmt.Println(m, len(m))

		//can be use
		delete(m, "ten")

		fmt.Println(m, len(m))
	*/

	//ex 29 map iteration
	/*
		m := map[int]string{1: "one", 2: "two", 3: "three", 4: "four"}

		for k, v := range m {
			fmt.Printf("key: %v | value: %v \n\r\n", k, v)
		}
	*/

	//ex 30 refs type map
	/*
		m := map[int]string{1: "one", 2: "two", 3: "three", 0: "zero"}

		fmt.Printf("map: %v | len: %v \n\r\n", m, len(m))
		chmap(m)
		fmt.Printf("map: %v | len: %v \n\r\n", m, len(m))
	*/
}

//ex 10 array are not linked type
/*
//foo func
func foo(a [4]int) {
	fmt.Println("inside function: ", a)
	a[2] = 555
	fmt.Println("inside function: ", a)
}
*/

//ex 21
/*
//ref func
func ref(sll []int) {
	if len(sll) > 0 {
		sll[0] = 333
	}
}
*/

//ex 30 ref type map
/*
//chmap func
func chmap(mm map[int]string) {
	mm[1] = "REFAAAYEEEEEEE!"
	delete(mm, 3)
}
*/
