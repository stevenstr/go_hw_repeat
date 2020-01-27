/*
 *Author: Stefan
 *Date: 01/27/2020
 *Last changes: 01/27/2020 21:49
 *Task: Class Work L3 Ex 1-9
 */

package main

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

}

/*
//ex 8 //work only via values
func chang(a [3]int) {
	a[1] = 777
	fmt.Println(a)
}
*/
