/*
 *Author: Stefan
 *Date: 01/27/2020
 *Last changes: 01/27/2020 01:35
 *Task: HW 2.2 Square
 */

package main

import "fmt"

//Point type
type Point struct {
	x, y int
}

//Square type
type Square struct {
	start Point
	a     uint
}

//End method
func (s Square) End() (p Point) {
	p.x = s.start.x + int(s.a)
	p.y = s.start.y + int(s.a)
	return p
}

//Perimeter method
func (s Square) Perimeter() (p uint) {
	p = s.a * uint(4)
	return p
}

//Area method
func (s Square) Area() (ar uint) {
	ar = s.a * s.a
	return ar
}

//main func
func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
