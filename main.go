/*
* Author: Stefan
* Created: 04.02.2020
* Last changes: 04.02.2020 21:09
* Task: Sample of basic golang servise, just implement basic feature
 */

package main

//import our deps
import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stevenstr/go_hw_repeat/ex1"
	"github.com/stevenstr/go_hw_repeat/ex2"
)

//main func
func main() {
	// create router using gorilla/mux
	r := mux.NewRouter()

	//set on "/hw" address the HelloBrouser function
	r.HandleFunc("/hw", ex1.HelloBrouserF)
	
	//set on "/hw" address the HelloBrouser function
	r.HandleFunc("/hw1", ex2.HelloBrouserS)

	//set http port for listen and serve on our router(mux wich was created)
	http.ListenAndServe(":8080", r)
}
