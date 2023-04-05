/**
package main

import "fmt"

func main() {
	//array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	//slice igual que un array pero no se le indica la longitud
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//metodos en el slice
	fmt.Println(slice[0])   // 0
	fmt.Println(slice[:3])  // 0,1,2 (posicion 3 es esclusiva)
	fmt.Println(slice[2:4]) // 2,3  (la posicion 4 es exclusiva)
	fmt.Println(slice[4:])  // 4,5,6

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nueva list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

}
