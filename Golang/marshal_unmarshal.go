package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name       string  `json:"name"`
	Age        int     `json:"age"`
	Percentage float64 `json:"perc"`
}

func main() {
	var p Person
	p = Person{Name: "Rida", Age: 26, Percentage: 85.98}
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}
	fmt.Println("Marshalled JSON:", string(data))

	var p2 Person
	err = json.Unmarshal(data, &p2)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}
	fmt.Println(p2.Name, p2.Age, p2.Percentage)
	fmt.Println("Unmarshalled JSON:", p2)
}
