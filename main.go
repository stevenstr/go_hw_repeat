/*
 *Author: Stefan
 *Date: 02/06/2020
 *Last changes: 02/06/2020 21:15
 *Task: CW L7
 */

package main

import (
	"encoding/json"
	"fmt"
)

//Pizza struct
type Pizza struct {
	Name     string
	Diameter int
	//unexported field
	//weight int
	Weight int
}

//serialization of struct in json format
var jsonStr = `{"name":"Cheese", "diameter": 35, "weight":500}`

func main() {
	//put jsonStr into slice byte
	data := []byte(jsonStr)

	p := &Pizza{}

	//Unmarshal our byte slice into p struct
	json.Unmarshal(data, p)
	//print our unmarshaled struct p
	fmt.Println(p)

	//Marshaling our struct into json
	pp, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Something wrong!")
	}
	fmt.Println(string(pp))
}
