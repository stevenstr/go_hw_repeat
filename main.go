/*
 *Author: Stefan
 *Date: 02/01/2020
 *Last changes: 02/01/2020 23:24
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
