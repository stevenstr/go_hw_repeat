/*
 *Author: Stefan
 *Date: 03/09/2020
 *Last changes: 03/09/2020 17:55
 *Task: Chapter 2 Lecture 2 HW 2.2 factorial
 */

package main

import "fmt"

//Point struct
type Point struct {
	x, y int
}

//Square struct
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
func (s Square) Perimeter() uint {
	return s.a * 4
}

//Area method
func (s Square) Area() uint {
	return s.a * s.a
}

//main func
func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())

	s1 := Square{Point{2, 1}, 6}
	fmt.Println(s1.End())
	fmt.Println(s1.Perimeter())
	fmt.Println(s1.Area())
}
