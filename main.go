/*
 *Author: Stefan
 *Date: 02/01/2020
 *Last changes: 02/05/2020 20:17
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
	//ex 13
	/*
		var (
			x  interface{}
			a  interface{} = [2]int{1, 2}
			sl interface{} = []int{1, 2}
		)
		fmt.Printf("DType:%T | value: %v\n\r\n", x, x)
		fmt.Printf("DType:%T | value: %v\n\r\n", a, a)
		fmt.Printf("DType:%T | value: %v\n\r\n", sl, sl)
	*/
	//ex 14
	/*
		n := Node{}
		n.Read()
		fmt.Println("Exec:")
		exec(n)

		var m Node //*Node - Panic because
		m.Read()
		fmt.Println("Exec:")
		exec(m)
	*/

	//Type assertion (extend interface)
	//ex 15
	/*
		//as struct
		var Pes = Dog{}
		Pes.voice()
		Pes.run()
		fmt.Printf("%T\n\r\n", Pes)

		//as interface, just input struct to interface
		var ppes Voicer = Dog{}
		ppes.voice()
		fmt.Printf("%T\n\r\n", ppes)

		//as Dog object, just output the struct
		pesik := ppes.(Dog)
		pesik.voice()
		pesik.run()
		fmt.Printf("%T\n\r\n", pesik)
	*/
	//ex 16
	/*
		var w io.Writer

		w = os.Stdout

		w.Write([]byte("sdfsdfsd\n"))
		//w.Read() //error

		rw := w.(io.ReadWriter)

		rw.Write([]byte("sdfsdfsd"))
		rw.Read([]byte("ok"))

		//t := w.(Animal) // panic
	*/

	//ex 17
	/*
		var w io.Writer
		w = os.Stdout

		rw, ok := w.(io.ReadWriter)
		fmt.Println(ok)
		fmt.Printf("%T\n\r\n", rw)

		r, ok := w.(Animal)
		//can't be extended becouse animal not consist
		fmt.Println(ok)
		fmt.Printf("%T\n\r\n", r)
	*/

	//ex 18
	/*
		var a interface{} = true    //bool
		var b interface{} = "foooo" //unknown
		var c interface{} = 4.4     //float64
		var d interface{}           //null

		fmt.Println("Dynamic Value: ")
		fmt.Println(dynamicValue(a))
		fmt.Println(dynamicValue(b))
		fmt.Println(dynamicValue(c))
		fmt.Println(dynamicValue(d))
	*/

	//Erroe handling
	//ex 19
	/*
		if _, err := fact(-5); err != nil {
			fmt.Println("Error: ", err)
		}

		if _, err := fact(0); err != nil {
			fmt.Println("Error: ", err)
		}

		if _, err := fact(5); err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
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

//ex 14
/*
//Node struct
type Node struct{}

//Read method
func (n Node) Read() {
	fmt.Println("READ!!!")
}

//Reader interface
type Reader interface {
	Read()
}

//exec func
func exec(r Reader) {
	fmt.Printf("%T\n\r\n", r)
	fmt.Println(r)
	if r != nil {
		r.Read()
	}
}
*/

//ex 15
/*
//Dog struct
type Dog struct{}

//voice method
func (d Dog) voice() {
	fmt.Println("Gav")
}

//run method
func (d Dog) run() {
	fmt.Println("Run")
}

//Voicer interface
type Voicer interface {
	voice()
}
*/

//ex 16
/*
type Animal interface {
	voice()
}
*/

//ex 17
/*
//Animal interface
type Animal interface {
	voice()
}
*/

//ex 18
/*
func dynamicValue(x interface{}) string {
	switch x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return "INT, UINT"
	case float64:
		return "FLOAT64"
	case bool:
		return "BOOL"
	default:
		return "Unknown"
	}
}
*/

//ex 19
/*
func fact(a int) (int, error) {
	if a < 0 {
		return 0, fmt.Errorf("Number is less then 0: ", a)
	}
	if a == 0 {
		return 1, nil
	}
	//...
	return 0, nil
}
*/
