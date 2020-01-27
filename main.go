/*
 *Author: Stefan
 *Date: 01/27/2020
 *Last changes: 01/27/2020 15:10
 *Task: Class Work L2 Ex 1-18
 */

package main

func main() {
	//ex 1
	/*
		var a int
		var name string = "hello world!"

		var (
			age  int
			g    string = "M"
			m, y        = 1, 2020
		)

		fmt.Println(a, name)
		fmt.Printf("a: %v %T | name: %v %T\n\r\n", a, a, name, name)
		fmt.Println(age, g, m, y)
		fmt.Printf("age: %v %T| gen: %v %T| m: %v %T| y: %v %T", age, age, g, g, m, m, y, y)
	*/
	//ex 2
	/*
		cat := "Murzik"
		cat, mouse := "Murzik", "Jerry"

		fmt.Println(cat, mouse)
	*/
	//ex 3
	/*
		a := 5
		if a < 0 || a > 4 {
			fmt.Println("foo")
		} else {
			fmt.Println("bar")
		}
	*/
	//ex 4
	/*
		//for style
		for i := 0; i <= 5; i++ {
			fmt.Println("i: ", i)
		}
		fmt.Printf("\n\r\n")
		//while style
		a := 0
		for a < 10 {
			fmt.Println("a: ", a)
			a++
		}
		fmt.Printf("\n\r\n")
		//do while style
		b := 0
		for {
			fmt.Println("b: ", b)
			if b == 2 {
				fmt.Println("Continue!")
				b++
				continue
			}
			if b == 3 {
				fmt.Println("!!!3!!!")
			}
			if b == 5 {
				fmt.Println("The end value is: ", b)
				break
			}
			b++
		}
	*/
	//ex 5
	/*
		var (
			a int
			b int32
			c int64
		)
		fmt.Println(a == int(b))
		fmt.Println(a == int(c))
	*/
	//ex 6
	/*
		a, b := 1, 1
		fmt.Println(a, b)
		a++
		b--
		fmt.Println(a, b)
		a += a
		b += a
		fmt.Println(a, b)
		a--
		b--
		fmt.Println(a, b)
		var c uint8 = 255
		fmt.Println(c)
		c++
		fmt.Println(c)
	*/
	//ex 7
	/*
		var app int32 = 1
		var orang int64 = 2
		result := int64(app) + orang
		fmt.Printf("app: %v %T | orang: %v %T | result: %v %T \n\r\n", app, app, orang, orang, result, result)
	*/
	//ex 8
	/*
		s := "Hello, МиР!"
		l := len(s)                     //return number of bytes, not always equals to num of symbols
		rn := utf8.RuneCountInString(s) // work with runes
		fmt.Println(l, rn)
		fmt.Println("Use len: ")
		for i := 0; i < l; i++ {
			fmt.Println(s[i])
		}
		fmt.Println()
		fmt.Println("Use runes: ")
		for i := 0; i < rn; i++ {
			fmt.Println(s[i])
		}
		fmt.Println()
		fmt.Println("Use range: ")
		for len(s) > 0 {
			r, size := utf8.DecodeRuneInString(s)
			fmt.Printf("%c %v\n", r, size)

			s = s[size:]
		}
	*/
	//ex 9
	/*
		type Minutes int

		var a int
		var b Minutes
		fmt.Printf("a: %T | b: %T \n\r", a, b)
		fmt.Println(a == int(b))
	*/
	//ex 10
	/*
		const a int = 55
		fmt.Printf("a: %v %T \n", a, a)

		const (
			b = 1
			c
			d = 2
			e
		)
		fmt.Println(b, c, d, e)
	*/
	//ex 11
	/*
		const (
			J int = iota
			F
			March
			A
			May
		)
		fmt.Println(J, F, March, A, May)
	*/
	//ex 12 factorial
	/*
		fmt.Println(fact(0))
		fmt.Println(fact(1))
		fmt.Println(fact(2))
		fmt.Println(fact(3))
		fmt.Println(fact(4))
		fmt.Println(fact(7))
	*/
	//ex 13
	/*
		fmt.Println(addm(2))
		fmt.Println(addm(3))
		fmt.Println(addm(4))
		fmt.Println(addm(5))
		fmt.Println(addm(6))
	*/

	//ex 14
	/*
		fmt.Println(even(2))
		fmt.Println(even(3))
		fmt.Println(even(5))
		fmt.Println(even(8))
		fmt.Println(even(23))
	*/

	//ex 15
	/*
		defer say()
		fmt.Println("Hello!")
		fmt.Println("Wait...")
	*/
	//ex 16
	/*
		for i := 0; i < 11; i++ {
			fmt.Println(i)
		}
		for i := 0; i < 11; i++ {
			defer fmt.Println(i)
		}
	*/
	//ex 17
	/*
		//HP struct
			type HP struct {
				CPU int
				SSD int
				PSU int
			}

			var hp255 HP
			hp255.CPU = 7
			hp255.SSD = 128
			fmt.Println(hp255)
			hp255.PSU = 350
			fmt.Println(hp255)
	*/
	//ex 18
	/*
		//Person struct
		type Person struct {
			id  int
			age int
		}

		var Rob Person
		fmt.Println(Rob)
		Rob = Person{1, 63}
		fmt.Println(Rob)

		var Ken Person
		fmt.Println(Ken)
		Ken = Person{id: 2}
		fmt.Println(Ken)
		Ken = Person{age: 76}
		fmt.Println(Ken)
		Ken = Person{id: 2, age: 76}
		fmt.Println(Ken)
	*/

}

//ex 12
/*
func fact(i uint) uint {
	if i == 0 {
		return 1
	}
	return i * fact(i-1)
}
*/
//ex 13
/*
func addm(i int) (int, int) {
	return i + i, i * i
}
*/
//ex 14
/*
func even(i int) (int, bool) {
	if i%2 == 0 {
		return i * 2, true
	} else {
		return i, false
	}
}
*/
//ex 15
/*
func say() {
	fmt.Println("Bye Bye!")
}
*/
