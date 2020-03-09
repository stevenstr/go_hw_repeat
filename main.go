/*
 *Author: Stefan
 *Date: 03/09/2020
 *Last changes: 03/09/2020 13:15
 *Task: Chapter 1 Lecture 1 Examples
 */

package main

import "fmt"

//const
const (
	j = iota
	k
	l
	m
)

//main func
func main() {
	//print some text
	fmt.Println("Hello Golang!")
	fmt.Printf("%s \n", "Hello world!")
	fmt.Println("Sould be new line")

	//some vars
	var a int
	var b int = 1
	var c = 2

	fmt.Printf("%d %d %d\n\r\n", a, b, c)

	//another way
	g := 9
	a = 11

	fmt.Println(g, a)

	//strings
	st := "Hello"
	st1 := " world!"

	fmt.Println(st, st1)

	//print consts
	fmt.Println(j, k, l, m)
}
