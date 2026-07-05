package main

import (
	"fmt"
)

func main() {
	// key - value
	// NY - NewYork
	m := map[string]string{
		"NY": "NewYork",
		"LA": "LosAngeles"}

	fmt.Println("Словарь городов:")
	fmt.Println(m)
	fmt.Println(m["LA"])
	m["LA"] = "Los Angeles"
	m["NY"] = "New York"
	fmt.Println(m["LA"])
	m["SF"] = "San Francisco"
	m["CH"] = "Chicago"
	fmt.Println(m)
	delete(m, "CH")
	fmt.Println(m)
}
