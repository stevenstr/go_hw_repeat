/*
* Author: Stefan
* Created: 04.02.2020
* Last changes: 04.02.2020 21:09
* Task: Sample of basic golang servise, just implement basic feature
 */

package ex1

import (
	"fmt"
	"net/http"
)

//HelloBrouserF func Return to brouser text using fmt package
func HelloBrouserF(w http.ResponseWriter, r *http.Request) {
	//set type of header for the http requests
	w.Header().Set("Content-Type", "application-json")
	fmt.Fprintf(w, "GOGOGOGO HELLO GOLANG AND IVAN, LESH AND KOLYA!!!!")
	fmt.Fprintf(w, "Using fmr.Fprinf()")
}
