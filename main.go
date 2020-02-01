/*
 *Author: Stefan
 *Date: 02/01/2020
 *Last changes: 02/02/2020 00:06
 *Task:Class Work L4 Ex
 */

package main

func main() {
	//ex 1
	/*
		d := dog{}
		d.voice()
	*/
	//ex 2
	/*
		var c ByteCounter
		c.Write([]byte("Hello"))
		fmt.Println(c)
		c = 0
		fmt.Println(c)
		fmt.Fprintf(&c, "hello world")
		fmt.Println(c)
	*/
	//ex 3
	/*
		var w io.Writer

		w = os.Stdout         // ok, os.File implements Write method
		w = new(bytes.Buffer) // ok, bytes.Buffer implements Write method
		//w = time.Second       // Error: time.Duration does not implement io.Writer (missing Write method)
	*/
	//ex 4
	/*
		//ok
		a := item{}
		a.A()
		a.B()
		//ok on pointer
		c := &item{}
		c.A()
		c.B()
		//ok
		item{}.A()
		//!ok in result of literal type are't addresseble
		//item{}.B()// can't be call on reciever
	*/
	//ex 5
	/*
		os.Stdout.Write([]byte("hello ok"))
		os.Stdout.Close()

		var w io.Writer //can call only Write method

		w.Write([]byte("can call write"))
		//w.Close() // mistace, interface writer does not implement Close method
	*/
	//ex 6
	/*
		var any interface{}

		any = "empty"
		fmt.Printf("%T %v\n\r\n", any, any)
		any = []byte("empty")
		fmt.Printf("%T %v\n\r\n", any, any)
		any = "hello empty"
		fmt.Printf("%T %v\n\r\n", any, any)
	*/
	//ex 7
	/*
		//should be ok
		var ccc animal = cat{}
		fmt.Println(ccc)
		ccc.voice()
	*/
	//ex 8
	/*
		//should be error - cat don't exist method voice
		var ccc animal = cat{}
		//fmt.Println(ccc)
		//ccc.voice()
	*/
	//ex 9 Media
	/*
		//Artifact interface
		//struct Album, Book, Song, Movie, Magazine
		type Artifact interface {
			Title() string
			Creators() string
			Created() time.Time
		}
		//Text interface
		//struct Book, Magazine
		type Text interface {
			Pages() uint
			Words() uint
		}
		//Audio interface
		//struct Album, Movie, Audio
		type Audio interface {
			Stream() (io.Reader, error)
			RunningTime() time.Duration
		}
	*/
	//ex 10
	/*
		var w io.Writer
		fmt.Printf("DType: %T | value: %v\n\r\n", w, w)
		w = os.Stdout //pointer to os.File
		fmt.Printf("DType:%T | value: %v\n\r\n", w, w)
		w = new(bytes.Buffer)
		fmt.Printf("DType:%T | value: %v\n\r\n", w, w)
		w = nil
		fmt.Printf("DType:%T | value: %v\n\r\n", w, w)
	*/
	//ex 11
	/*
		var x interface{} = []int{1, 2, 3}
		//type x as var - interface
		//dinamic type - []int - slice
		//dinamic value = []int{1,2,3} - slice
		fmt.Printf("DType:%T | value: %v\n\r\n", x, x)
	*/
	//ex 12
	/*
		var x interface{} = [3]int{1, 2, 3}
		var y interface{} = [3]int{1, 2, 3}
		var z interface{} = [3]string{""}
		var n interface{} = [3]int{3, 4, 5}

		fmt.Printf("DType:%T | value: %v\n\r\n", x, x)
		fmt.Printf("DType:%T | value: %v\n\r\n", y, y)
		fmt.Printf("DType:%T | value: %v\n\r\n", z, z)
		fmt.Printf("DType:%T | value: %v\n\r\n", n, n)

		fmt.Println("x == y", x == y) //true
		fmt.Println("x == n", x == n) //false - different values
		fmt.Println("x == z", x == z) //false - different types
	*/

}

//ex 1
/*
//animal interface
type animal interface {
	voice()
}

//dog struct
type dog struct{}

//voice method
func (_ dog) voice() {
	fmt.Println("Gav")
}
*/

//ex 2
/*
//ByteCounter type
type ByteCounter int

//Write method
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
*/
//ex 3
/*
//item struct
type item struct{}

//A method
func (i item) A() {
	fmt.Println("A")
}

//B method
func (i *item) B() {
	fmt.Println("B")
}
*/

//ex 7
/*
//cat struct
type cat struct{}

//voice method
func (_ cat) voice() {
	fmt.Println("Myau!")
}

//animal interface
type animal interface {
	voice()
}
*/

//ex 8
/*
//cat struct
type cat struct{}

//animal interface
type animal interface {
	voice()
}
*/
