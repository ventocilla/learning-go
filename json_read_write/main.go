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
			"has_dog": true,
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false,
		}
	]
	`

	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		fmt.Println("Error unmarshalling json", err)
	}
	fmt.Printf("unMarsheled: %v", unmarshalled)
}
