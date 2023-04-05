/**
package main

import (
	"fmt"
	"log"
	"strconv"
)

//atoi
//Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.

//fatal
//Fatal is equivalent to Print() followed by a call to os.Exit(1).

func esPar(num int) string {
	if num%2 == 0 {
		return "Es Par"
	} else {
		return "No es Par"
	}
}

func esUsuario(usuario string, password string) string {
	if usuario == "tatiana" && password == "1234" {
		return "Es usuario"
	} else {
		return "No es usuario, usuario y/o contrasena incorrecta"
	}

}

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// AND
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	// OR
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad, OR")
	}

	//Convertir texto a numero
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)

	esNumPar := esPar(3)
	fmt.Println("par: ", esNumPar)

	esUsuarioGo := esUsuario("tatiana", "1234")
	fmt.Println("Es usuario: ", esUsuarioGo)
}
