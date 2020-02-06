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

//main func
func main() {
	//ex 1
	/*
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
	*/

	//ex 2

	p := &Pizza{
		Name:  "Cheese",
		Price: 15,
	}

	//Marshaling struct into json
	result, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Somethong wrong!")
	}
	fmt.Println(string(result))

	//Unmarshaling json as slice byte into struct
	res := []byte(result)
	c := &Pizza{}
	json.Unmarshal(res, c)
	fmt.Println(c)
}

//ex 1
/*
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
*/

//ex 2
type Pizza struct {
	Name  string `json:"pizza_name"`
	Price int    `json:"price"`
}
