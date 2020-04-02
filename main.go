/*
* Author: Stefan
* Created: 04.02.2020
* Last changes: 04.03.2020 00:17
* Task: Lesson 2 Class Work Vars, arrays, structure, methods
 */

package main

//import our deps

//main func
func main() {
	//ex 1 Write hello world
	/*
		fmt.Println("Hello golang world!!!")
	*/

	//ex 2 Use packages
	//import "fmt"
	//import "github.com/stevenstr/go_hw_repeat/laguage"
	/*
		fmt.Println(laguage.Language)
		//fmt.Println(laguage.foo) //error(unexported variable)

	*/

	//ex 3 Variables
	/*
		var a int
		var b int = 33
		var c = 33

		d := 666

		fmt.Println(a, b, c, d)
		fmt.Printf("a: %5T | b: %5T | c: %5T | d: %5T \r\n\r", a, b, c, d)

		f := true
		g := 3.14

		fmt.Println(f, g)
		fmt.Printf("f: %5T | g: %5T \r\n\r", f, g)

		s := "Hello"
		var st string
		fmt.Printf("%s | %s \r\n\r", s, st)
		fmt.Printf("%T | %T \r\n\r", s, st)
	*/

	//ex 4 Group of vars
	/*
		var (
			a int
			b int = 333
			c     = "Kolya"
			d     = ""
		)

		fmt.Printf("%5v | %5T \r\n\r", a, a)
		fmt.Printf("%5v | %5T \r\n\r", b, b)
		fmt.Printf("%5v | %5T \r\n\r", c, c)
		fmt.Printf("%5v | %5T \r\n\r", d, d)
	*/

	//ex 5 Short var decloration (need fix)
	/*
		age := 22
		cat := "Murzik"
		i, mouse := 1, "Jerry"
		// no new variables on left side of :=
		cat = "Tom"
		// no new variables on left side of :=
		cat, mouse = "Tom", "NoName"
		// at least one new variable
		cat, newVar := "Tom", 1
		fmt.Println(age, cat, i, mouse, newVar)
	*/

	//ex 6 Condition if-else
	/*
		a := 5

		if a == 0 || a > 4 {
			fmt.Println("Bigger then 4 or equal zero!")
		} else {
			fmt.Println("Other cases....")
		}
	*/

	//ex 7 Loops (for)
	/*
		for i := 0; i < 6; i++ {
			fmt.Println(i)
		}
		fmt.Printf("\n\r\n\r")

		a := 8

		for a > 0 {
			fmt.Println(a)
			a--
		}
		fmt.Printf("\n\r\n\r")

		a = 0

		for {
			fmt.Println(a)
			if a == 7 {
				break
			}
			fmt.Println(a)
			a++
		}
		fmt.Printf("\n\r\n\r")
	*/

	//ex 8 Type compare
	/*
		var (
			a int32
			b int64
		)

		//type mismatched
		//fmt.Println(a == b)
		//should be OK
		fmt.Println(int64(a) == b)
	*/

	//ex 9 Integer operator
	/*
		a := 0
		b := 4
		fmt.Println(a, b)

		a += b
		b += a
		fmt.Println(a, b)

		a--
		b++
		fmt.Println(a, b)

		b--
		a++
		fmt.Println(a, b)

		fmt.Println(19 % 2)
	*/

	//ex 10 Conver types
	/*
		var (
			apples int32 = 3
			orange int64 = 2
		)
		fmt.Printf("%5T | %5T \n\r\n", apples, orange)

		var compote int = int(apples) + int(orange)

		fmt.Printf("compote: %5v | Type: %5T \r\n\r", compote, compote)
	*/

	//ex 11 Strings
	/*
		s := "Hello, world!"

		fmt.Printf("%v | %T \r\n\r", s, s)

		hello := s[:5]
		world := s[7:]

		fmt.Printf("%v | %T", hello, hello)
		fmt.Printf("%v | %T", world, world)

		fmt.Println(s, hello, world)
	*/

	//ex 12 bytes
	/*
		s := "Приве"
		se := "Hello"

		//return that russian sympol take 2 byte in memory
		fmt.Printf("%v | %v \r\n\r", s, len(s))
		for i, v := range s {
			fmt.Println(i, v, s[i])
		}

			for len(s) > 0 {
				r, size := utf8.DecodeRuneInString(s)
				fmt.Println("size: ", size)
				fmt.Printf("%c %v\n", r, size)
				s = s[size:]
			}

			//return that eng symbol take 1 byte in memory
			fmt.Printf("%v | %v \r\n\r", se, len(se))

			for len(se) > 0 {
				r, size := utf8.DecodeRuneInString(se)
				fmt.Printf("%c %v\n", r, size)
				se = se[size:]
			}

			//it is mean that so that running on the string we
			//should know how much bytes our symbols take in memory
			// or sould use special type rune
	*/
	//ex 13 Custum types
	/*
		type Min int8
		type Max int64

		var a Min
		var b Max

		fmt.Println(a, b)
		fmt.Printf("type A: %T\r\n\r", a)
		fmt.Printf("type B: %T\r\n\r", b)

		//error
		//fmt.Println(a == b)

		//should be OK
		fmt.Println(Max(a) == b)
	*/

	//ex 14 Constant, iota, features
	/*
		const (
			a = 1
			b
			c = 33
			d
		)

		fmt.Println(a, b, c, d)

		const (
			jan = iota
			feb
			mar
			apr
		)

		fmt.Println(jan, feb, mar, apr)
	*/

	//ex 15 functions
	/*
		a, b := 6, 2

		fmt.Println(add(a, b))
		fmt.Println(sub(a, b))
		fmt.Println(mult(a, b))
		fmt.Println(devi(a, b))
	*/

	//ex 16 Recursion
	/*
		var a uint = 5
		fmt.Println(rec(a))
	*/

	//ex 17 Multiple return value
	/*
		res1, res2 := jjjetVal(5, 5)
		fmt.Println(res1, res2)
	*/

	//ex 18 Multiple return value impl
	/*
		var (
			sum int
			ok  bool
		)

		if sum, ok = increaseEven(6); ok {
			fmt.Println("Even!!!")
		} else {
			fmt.Println("Odd!!!")
		}

		fmt.Println(sum)
	*/

	//ex 19 COOL DEFEEEEEERRRR!	defer add
	/*
		fmt.Println("FHFHSKDJFHSDKJFHSDKJF")
		defer add()
	*/

	//ex 20 Defer
	/*
		for i := 0; i < 5; i++ {
			defer fmt.Println(i)
			fmt.Println(i)
		}
	*/

	//ex 21 Structures
	/*
		type Book struct {
			ID     int
			Title  string
			Author string
			Pages  uint
		}

		//full form, each fields sould be passed
		Harry := Book{1, "Harry potter", "Lesha", 554}
		Rings := Book{2, "Lord of the rings", "Tolkin", 1000}

		//short form, we can passed only main for us fields (not each)
		Warcraft := Book{ID: 3, Title: "WOW Garrosh", Pages: 414}

		fmt.Println(Harry)
		fmt.Println()
		fmt.Println(Rings)
		fmt.Println()
		fmt.Println(Warcraft)
	*/

	//ex 22
	/*
		type Book struct {
			ID     int
			Title  string
			Author string
			Pages  uint
		}

		var a Book

		fmt.Println(a)
		fmt.Println(a.ID)
		fmt.Println(a.Author)

		a.ID = 1
		a.Author = "Kolya"
		a.Pages = 333
		a.Title = "How to become DOCKER MASTER"

		fmt.Println(a)
		fmt.Println(a.ID)
		fmt.Println(a.Author)
	*/

	//ex 23 Structure embedding
	/*
		type Engine struct {
			Power int
			Year  int
		}
		type Car struct {
			engine Engine
			Title  string
			Year   int
		}

		audi := Car{engine: Engine{Power: 70, Year: 1980}, Title: "Benz", Year: 2000}
		bmw := Car{Engine{67, 1985}, "BMW", 1998}
		opel := Car{Engine{88, 1999}, "OPEL", 1999}

		bmw.engine.Power = 95

		fmt.Println(audi, bmw, opel)
	*/

	//ex 24 Anonimous fields
	/*
		type Engine struct {
			Power int
			Year  int
		}

		type Car struct {
			Engine
			Year int
		}

		benz := Car{Engine{550, 2016}, 2019}
		fmt.Println(benz)

		benz.Power = 633
		fmt.Println(benz)

		fmt.Println()

		//full form return 2016
		fmt.Println(benz.Engine.Year)
		//short form return 2019 (not dive onto the embedding field)
		fmt.Println(benz.Year)
	*/

	//ex 25 Methods

	/*
		p := Point{1, 1}
		q := Point{3, 3}

		//run method on our p structure and send her as parameter q structure
		//work with value because p it is aren't pointer
		fmt.Println(p.Dist(q))
	*/

	//ex 26  Method with pointer reciever
	/*
		p := Point{1, 1}

		p.Scale1()
		fmt.Println(p) // 1  1

		p.Scale2()     // it is changed source because using pointer reciever
		fmt.Println(p) // and return is 2 2
	*/
}

//#################################################
//#########    Function to exercises    ###########
//#################################################

//ex 5
//error, short form can be used only in function
//age := 22

//ex 15
/*
func add(x, y int) int  { return x + y }
func sub(x int, y int) int  { return x - y }
func mult(x, y int) int { return x * y }
func devi(x, y int) int { return x / y }
*/

//ex 16 Recursion
/*
func rec(a uint) uint {
	if a == 0 {
		return 1
	}
	return a * rec(a-1)
}
*/

//ex 17 Multiple return value
/*
func jjjetVal(a, b int) (int, int) {
	return a + 5, b - 5
}
*/

//ex 18
/*
func increaseEven(a int) (int, bool) {
	if a%2 == 0 {
		return a * 2, true
	} else {
		return a, false
	}
}
*/

//ex 19 COOL DEFEEERRRR!!!
/*
func add() {
	fmt.Println(566666666 + 1)
}
*/

//ex 25 Methods
/*
//Point struct
type Point struct {
	x float64
	y float64
}

//Dist method
func (p Point) Dist(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}
*/

//ex 26 Method with pointer reciever
/*
//Point struct
type Point struct{ X, Y int }

//Scale1 method
func (p Point) Scale1() {
	p.X *= 2
	p.Y *= 2
}

//Scale2 method with pointer reciever
func (p *Point) Scale2() {
	p.X *= 2
	p.Y *= 2
}
*/
