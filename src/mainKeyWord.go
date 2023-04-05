/**
package main

import "fmt"

func main() {
	//Defer es la ultima linea que se va a ejecutar
	defer fmt.Println("")
	fmt.Println("Mundo")

	//Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}
		//Break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}
