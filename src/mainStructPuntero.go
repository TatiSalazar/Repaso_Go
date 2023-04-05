/**
package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.disk, "pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {
	a := 50
	b := &a // El “&” accede a la dirección del espacio de memoria de la variable.

	fmt.Println(b)
	fmt.Println(*b) // "*" accede al valor que contiene ese espacio de memoria,
	//dado el nombre de una variable o una dirección especifica.

	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)

}
**/