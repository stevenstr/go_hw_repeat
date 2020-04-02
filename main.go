/*
* Author: Stefan
* Created: 04.02.2020
* Last changes: 04.02.2020 20:56
* Task: Sample of basic golang servise, just implement basic feature
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//HelloBrouser func
func HelloBrouser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application-json")

	sl := []string{"Hello world!"}

	json.NewEncoder(w).Encode(sl)
	fmt.Fprintf(w, "GOGOGOGO HELLO GOLANG AND IVAN, LESH AND KOLYA!!!!")
}

//main func
func main() {
	// create router using gorilla/mux
	r := mux.NewRouter()

	//set on "/hw" address the HelloBrouser function
	r.HandleFunc("/hw", HelloBrouser)

	//set http port for listen and serve on our router(mux wich was created)
	http.ListenAndServe(":8080", r)
}
