/*
 *Author: Stefan
 *Date: 02/05/2020
 *Last changes: 02/05/2020 20:39
 *Task: HW 4.2
 */

package main

import "fmt"

//Pi const
const Pi float64 = 3.14159265359

//Figure interface
type Figure interface {
	area() float64
	perimeter() float64
}

//Square struct
type Square struct {
	a float64
}

//Methods for square
//area square method
func (s Square) area() float64 {
	return s.a * s.a
}

//perimeter square method
func (s Square) perimeter() float64 {
	return s.a * 4.0
}

//Circle struct
type Circle struct {
	r float64
}

//Methods for Circle
//area circle method
func (c Circle) area() float64 {
	return Pi * (c.r * c.r)
}

//perimeter circle method
func (c Circle) perimeter() float64 {
	return 2.0 * Pi * c.r
}

//main function
func main() {
	var s Figure = Square{2}
	var c Figure = Circle{2}

	fmt.Println(s.area(), s.perimeter())
	fmt.Println(c.area(), c.perimeter())
}
