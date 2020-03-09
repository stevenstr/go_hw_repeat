/*
 *Author: Stefan
 *Date: 03/09/2020
 *Last changes: 03/09/2020 17:18
 *Task: Chapter 2 Lecture 2 Basics, struct, functions, methods
 */

package main

//"github.com/stevenstr/go_hw_repeat/bye"
//"github.com/stevenstr/go_hw_repeat/hi"
//"github.com/stevenstr/go_hw_repeat/person"

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

	//ex 15 String
	/*
		s := "Hello, world!"
		fmt.Println(s)

		//mistake! String are immutable
		//s[3] = "g" //mistake! String are immutable

		s1 := s[:5]
		s2 := s[5:]
		fmt.Println(s1, s2)

		//concatenation!
		s3 := s1 + s2
		fmt.Println(s3)

		s4 := s[:6] + " " + s[7:]
		fmt.Println(s4)
	*/

	//ex 16 Work with runes
	/*
		arrEng := [...]string{"W", "o", "r", "l", "d", "!"}
		arrRus := [...]string{"М", "и", "р", "н", "ы", "й"}

		fmt.Println("ENG")
		for _, v := range arrEng {
			fmt.Println(v, "| Runes: ", utf8.RuneCountInString(v), "| len: ", len(v))
		}

		fmt.Println()

		fmt.Println("RUS")
		for _, v := range arrRus {
			fmt.Println(v, "| Runes: ", utf8.RuneCountInString(v), "| len: ", len(v))
		}
	*/

	//ex 17
	/*
			s := "Привет World!"

			rn := utf8.RuneCountInString(s)
			ln := len(s)

			fmt.Println(s)
			fmt.Println("Rune: ", rn)
			fmt.Println("Len: ", ln)

			fmt.Println("Using for Range:")
			for _, r := range s {
				fmt.Printf("%c \n", r)
			}
			fmt.Println()

			fmt.Println()
			for i, v := range s {
			fmt.Printf("%d | %v | %c \n\r\n", i, v, v)

		}
	*/

	//ex 18 Custom type
	/*
		type Minutes int32
		type Hours int32

		var a Minutes = 21
		var b Hours = 21

		fmt.Println(a == Minutes(b))
		fmt.Println(int(a) == int(b))
	*/

	//ex 20 iota const
	/*
		const (
			a = 1.1
			b
			c = 2.2
			d
		)

		fmt.Printf("%v %T\n\r\n", a, a)
		fmt.Printf("%v %T\n\r\n", b, b)
		fmt.Printf("%v %T\n\r\n", c, c)
		fmt.Printf("%v %T\n\r\n", d, d)
		fmt.Println(a, b, c, d)
		fmt.Println()

		const (
			J = iota
			F
			M
		)

		fmt.Println(J, F, M)
		fmt.Printf("%v %T\n\r\n", J, J)
		fmt.Printf("%v %T\n\r\n", F, F)
		fmt.Printf("%v %T\n\r\n", M, M)
		fmt.Println()
	*/

	//ex 21 functions
	/*
		fmt.Println(add(2, 3))
		fmt.Println(sub(5, 3))
		fmt.Printf("%T \n\r\n", add)
		fmt.Printf("%T \n\r\n", sub)
	*/

	//ex 22 factorial using recursion
	/*
		fmt.Println(factorial(5))
		fmt.Println(factorial(0))
		fmt.Println(factorial(13))
		fmt.Println(factorial(7))
	*/

	//ex 23 getVal
	/*
		fmt.Println(getVal(3, 3))
		fmt.Println(getVal(7, 3))
		fmt.Println(getVal(2, 8))
		fmt.Println(getVal(9, 5))
		fmt.Println(getVal(0, 3))
	*/

	//ex 24 Even Checker
	/*
		var (
			sum int
			ok  bool
		)

		if sum, ok = evenChecker(6); ok {
			fmt.Println("Even !!!")
		} else {
			fmt.Println("ODD!!!")
		}
		fmt.Println(sum)
	*/

	//ex 25 Defer statement
	/*
		defer fmt.Println("close connection...bye!")
		fmt.Println("Hello!")
		fmt.Println("do something...")
	*/

	//ex 26
	/*
		fmt.Println("Without defer: ")
		for i := 0; i < 6; i++ {
			fmt.Println(i)
		}

		fmt.Println("Using defer: ")
		for i := 0; i < 6; i++ {
			defer fmt.Println(i)
		}
	*/

	//ex 27 Structure
	/*
		//HarryPotter struct
		type HarryPotter struct {
			ID    uint
			Page  uint
			Title string
		}

		var harry HarryPotter
		harry.ID = 1
		harry.Title = "Prince"
		harry.Page = 455

		fmt.Println(harry)
		fmt.Printf("%v | %T \n", harry, harry)

		harry.ID++
		fmt.Printf("%v | %T \n", harry, harry)

		var mort HarryPotter
		fmt.Println(mort.ID)
		fmt.Println(mort.Page)
		fmt.Println(mort.Title)
		fmt.Println(mort)
	*/

	//ex 28 Person Exported/unexported fields
	/*
		var a person.Person

		a.Name = "Kolya"
		a.Age = 27
		//error
		//a.id = 1
		fmt.Println(a)
	*/

	//ex 29 Struct literals
	/*
		type Guse struct {
			ID     uint
			Legs   int
			Colour string
		}

		myGoose := Guse{1, 2, "white"}
		myGoose1 := Guse{Legs: 2, Colour: "yellow"} //ID default 0
		myGoose2 := Guse{Legs: 2, ID: 2, Colour: "black"}
		myGoose3 := Guse{Legs: 2, ID: 4, Colour: "green"}

		fmt.Println(myGoose)
		fmt.Println(myGoose1)
		fmt.Println(myGoose2)
		fmt.Println(myGoose3)
	*/

	//ex 30 Struct embedding
	/*
		//Engine struct
		type Engine struct {
			Power int
			Year  int
		}
		//Car struct
		type Car struct {
			engine Engine
			Title  string
		}

		myCar := Car{Engine{72, 1980}, "w123"}
		myCar1 := Car{Engine{80, 1999}, "VAZ2104"}
		myCar2 := Car{Engine{70, 1993}, "VW Golf 3"}
		myCar3 := Car{engine: Engine{Power: 78, Year: 1996}, Title: "Civic"}

		fmt.Println(myCar)
		fmt.Println(myCar1)
		fmt.Println(myCar2)
		fmt.Println(myCar3)
	*/

	//ex 31

}

//ex 21
/*
//add func
func add(x, y int) int {
	return x + y
}

//sub func
func sub(x, y int) (z int) {
	z = x - y
	return
}
*/

//ex 22
/*
//factorial func
func factorial(i uint) uint {
	if i == 0 {
		return 1
	}
	return i * factorial(i-1)
}
*/

//ex 23
/*
//getVal func
func getVal(x, y int) (int, int) {
	return x + y, x - y
}
*/

//ex 24
/*
//evenChecker func
func evenChecker(a int) (int, bool) {
	if a%2 == 0 {
		return a * 2, true
	}
	return a, false
}
*/
