package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false
		}
	]
	`

	// unMarshal => Json to Struct
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		fmt.Println("Error unmarshalling json", err)
	}
	fmt.Printf("unMarsheled: %v", unmarshalled)

	// Write json from a struct, Marshal => struct to Json
	var mySlice []Person
	var m1 Person
	var m2 Person

	m1.Firstname = "Wally"
	m1.Lastname = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	m2.Firstname = "Diana"
	m2.Lastname = "Prince"
	m2.HairColor = "black"
	m2.HasDog = false

	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)
	newJson, err := json.MarshalIndent(mySlice, "", "     ")
	if err != nil {
		fmt.Println("error marshalling", err)
	}
	fmt.Println(string(newJson))
}
