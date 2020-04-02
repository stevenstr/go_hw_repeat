/*
* Author: Stefan
* Created: 04.02.2020
* Last changes: 04.02.2020 21:11
* Task: Sample of basic golang servise, just implement basic feature
 */

package ex2

import (
	"encoding/json"
	"net/http"
)

//HelloBrouserS func Using json.NewEncoder to send some text to brouser
func HelloBrouserS(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application-json")

	sl := []string{"Hello world!", "Using json Encoder"}

	json.NewEncoder(w).Encode(sl)
}
