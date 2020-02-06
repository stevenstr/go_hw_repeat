/*
 *Author: Stefan
 *Date: 02/06/2020
 *Last changes: 02/06/2020 18:57
 *Task: HW 7.1
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//query struct
type qwery struct {
	Host      string `json:"host"`
	UA        string `json:"user_agent"`
	Ruri      string `json: "request_uri"`
	Accept    string `json:"Accept"`
	UserAgent string `json:"User-Agent"`
}

//GetData func
func GetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := qwery{
		Host:      r.Host,
		UA:        r.UserAgent(),
		Ruri:      r.RequestURI,
		Accept:    r.Header.Get("Accept"),
		UserAgent: r.Header.Get("User-agent"),
	}

	//pack for Write(interface)
	pp, err := json.Marshal(&p)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(pp)
	return
}

//main func
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", GetData)

	fmt.Println("starting server at: 8080")
	http.ListenAndServe(":8000", r)
}
