package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

)

type( 
	Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
    People struct {
	Peoples []Person `json:"peoples"`
}
)

// Create a variable to hold the decoded data
var people People

func main() {
	file, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode JSON data from the file into the People struct
	err = json.NewDecoder(file).Decode(&people)
	if err != nil {
		log.Fatal(err)
	}

	for i, person := range people.Peoples {
		fmt.Printf("Person %d:\n", i+1)
		fmt.Println("Name:", person.Name)
		fmt.Println("Age:", person.Age)
		fmt.Println()
	}
}
