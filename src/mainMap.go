/**
package main

import "fmt"

func main() {
	m := make(map[string]int) //tipo de dato string

	m["jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer un map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar valor
	value, ok := m["jose"]
	fmt.Println(value, ok)
}
