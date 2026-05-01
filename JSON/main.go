package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"personName"` // struct tag
	Age  int    `json:"-"`
	City string `json:"city"`
}

func main() {
	// var p person
	// p = person{
	// 	Name: "John",
	// 	Age:  25,
	// 	City: "Dhaka",
	// }

	// rawJson, err := json.Marshal(p)

	// if err != nil {
	// 	fmt.Println("Error", err)
	// }
	// fmt.Println(string(rawJson))

	var p2 person
	jsonText := `{"personName":"John","city":"Dhaka"}`

	error := json.Unmarshal([]byte(jsonText), &p2)

	if error != nil {
		fmt.Println("Error", error)
	}

	fmt.Println(p2.Name)
	fmt.Println(p2.Age)
	fmt.Println(p2.City)
}
