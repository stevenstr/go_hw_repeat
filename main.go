/*
 *Author: Stefan
 *Date: 03/09/2020
 *Last changes: 03/09/2020 14:45
 *Task: Chapter 2 Lecture 2 Basics, struct, functions, methods
 */

package main

//"github.com/stevenstr/go_hw_repeat/bye"
//"github.com/stevenstr/go_hw_repeat/hi"

//main func
func main() {
	// ex 1
	//fmt.Println("Hello Golang World!")

	//ex 2
	/*
		fmt.Println("Text:")
		hi.Hi()
		fmt.Println(bye.Name)

		//return mistake
		//fmt.Println(bye.name)
	*/

	//ex 3 value of vars
	/*
		var foo1 int = 55
		var foo2 = 4.3
		var foo3 int

		fmt.Println(foo1, foo2, foo3)

		bar1 := 2
		bar2 := "bar"
		bar3 := true
		bar4 := 3.33333
		bar5 := &bar1
		var bar6 *int
		var bar7 string
		var bar8 int
		var bar9 bool

		fmt.Printf("v: %8v | type: %8T \n\r\n", bar1, bar1)
		fmt.Printf("v: %8v | type: %8T \n\r\n", bar2, bar2)
		fmt.Printf("v: %8v | type: %8T \n\r\n", bar3, bar3)
		fmt.Printf("v: %8v | type: %8T \n\r\n", bar4, bar4)
		fmt.Printf("v: %8v | type: %8T \n\r\n", bar5, bar5)
		fmt.Printf("v: %8v | type: %8T \n\r\n", *bar5, *bar5)
		fmt.Printf("v: %8v | type: %8T \n\r\n", bar6, bar6)
		fmt.Printf("v: %8v | type: %8T \n\r\n", bar7, bar7)
		fmt.Printf("v: %8v | type: %8T \n\r\n", bar8, bar8)
		fmt.Printf("v: %8v | type: %8T \n\r\n", bar9, bar9)
	*/

	//ex 4 Var dec example
	/*
		var name string = "Ivan"
		var name1 string = "Lesha"

		var (
			a int
			b int = 4

			c, d = 777, "Azino"
		)

		fmt.Println(name, name1, a, b, c, d)

		fmt.Printf("val: %5v | type: %8T \n\r\n", name, name)
		fmt.Printf("val: %5v | type: %8T \n\r\n", name1, name1)
		fmt.Printf("val: %5v | type: %8T \n\r\n", a, a)
		fmt.Printf("val: %5v | type: %8T \n\r\n", b, b)
		fmt.Printf("val: %5v | type: %8T \n\r\n", c, c)
		fmt.Printf("val: %5v | type: %8T \n\r\n", d, d)
	*/

	//ex 5
	/*
		cat := "Murzik"
		i, mouse := 1, "Jerry"
		//error, no exist new variables
		//cat, mouse := "Jo", "Kuku"

		fmt.Println(cat, i, mouse)
		//ok
		a, mouse := 2, "BIG JERRY"

		fmt.Println(cat, i, a, mouse)
	*/

	//ex 6 if-else
	/*
		var foo int
		fmt.Scanf("%d", &foo)
		if foo == 1 {
			fmt.Println("Your input is 1)")
		} else if foo == 2 {
			fmt.Println("Your input is 2))")
		} else if foo == 3 {
			fmt.Println("Your input is 3)))")
		} else {
			fmt.Println("Bigger then 3 or less then 1!")
		}
	*/

	//ex 7
	/*
		a := 5

		if a < 0 || a > 4 {
			fmt.Println("foo")
		} else {
			fmt.Println("bar")
		}
	*/

	//ex 8 LOOPS
	/*
		//classic for style
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}

		fmt.Println()

		a := 10
		//classic while style
		for a >= 0 {
			fmt.Println(a)
			a--
		}

		fmt.Println()

		a1 := 6
		//while-forever style
		for {
			if a1 == 4 {
				fmt.Println("continue for a1 = ", a1)
				a1--
				continue
			}
			if a1 == 2 {
				fmt.Println("continue for a1 = ", a1)
				a1--
				continue
			}
			if a1 == 0 {
				break
			}
			fmt.Println(a1)
			a1--
		}
		fmt.Println("Loop is stay when a1 = ", a1)
		fmt.Println()

		//for-range style
		arr := [...]int{1, 2, 3, 4, 5, 6}

		for i, v := range arr {
			fmt.Println("index: ", i, "value: ", v)
		}
	*/

	//ex 9 int compate
	/*
		var (
			a int
			b int16
			c int64
		)

		//fmt.Println(a == b) //false
		//fmt.Println(b == c) //false
		//fmt.Println(a == c) //false

		fmt.Println(a == int(b)) //true
		fmt.Println(a == int(c)) //true
		fmt.Println(int(c) == int(b)) //true
	*/

	//ex 10
	/*
		a := 1

		const (
			Big   = 1 << 4
			Small = Big >> 2
		)

		fmt.Println(a)
		fmt.Println("Big: ", Big)
		fmt.Println("Small: ", Small)
	*/

	//ex 11
	/*
		a := 1
		a += 5

		fmt.Printf("%d\n\r\n", a) //decimal
		fmt.Printf("%b\n\r\n", a) //binary

		a--
		fmt.Println(a)
		a++
		fmt.Println(a)

		var b uint8 = 255
		fmt.Println(b)
		b += 1
		fmt.Println(b)
	*/

	//ex 12 Convert types
	/*
		var apples int32 = 1
		var oranges int8 = 3

		//fmt.Println(apples == oranges) //compile error
		fmt.Println(apples == int32(oranges)) //ok

		var compotik = apples + int32(oranges)
		fmt.Printf("value: %6v | type: %6T  \n\r\n", apples, apples)
		fmt.Printf("value: %6v | type: %6T  \n\r\n", oranges, oranges)
		fmt.Printf("value: %6v | type: %6T  \n\r\n", compotik, compotik)
	*/
	//ex 13
	/*
		t := true
		f := false

		fmt.Println("         ||")
		fmt.Println(t, "||", f, "===", (t || f))
		fmt.Println(t, "||", t, "===", (t || true))
		fmt.Println(f, "||", f, "===", (f || false))
		fmt.Println(f, "||", t, "===", (f || t))

		fmt.Println("         &&")
		fmt.Println(t, "&&", f, "===", (t && f))
		fmt.Println(t, "&&", t, "===", (t && true))
		fmt.Println(f, "&&", f, "===", (f && false))
		fmt.Println(f, "&&", t, "===", (f && t))

		fmt.Println("         !")
		fmt.Println("!", f, "===", (!f))
		fmt.Println("!", t, "===", (!t))
	*/

	//ex 14
	/*
		var a byte = 1
		var mask byte = 101

		var xora byte = ^a
		var anda byte = a & mask
		var xoranda byte = a &^ mask
		var ora byte = a | mask

		fmt.Printf("%b \n\r\n", a)
		fmt.Printf("%b \n\r\n", xora)
		fmt.Printf("%b \n\r\n", anda)
		fmt.Printf("%b \n\r\n", xoranda)
		fmt.Printf("%b \n\r\n", ora)
	*/

}
