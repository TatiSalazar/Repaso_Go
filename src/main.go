/**
package main

import (
	"fmt"
	"math"
)

func main() {

	// Paquete fmt para imprimir en consola
	helloMessage := "Hello"
	worldMessage := "world"

	// Println: Salto de Linea Automatico
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	// Con valores seguros
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	// Con valores inseguros
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)
	/*
	 bool:                    %t
	 int, int8 etc.:          %d
	 uint, uint8 etc.:        %d, %#x if printed with %#v
	 float32, complex64, etc: %g
	 string:                  %s
	 chan:                    %p
	 pointer:                 %p
	*/

	// Sprintf lo guarda como un string
	message := fmt.Sprintf("%v tiene más de %v cursos\n", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos: te imprime el tipo de dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

	// Declaracion de constantes
	const pi1 float64 = 3.14
	const pi2 = 3.14
	fmt.Println(pi1, pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("area cuadrado: ", areaCuadrado)

	// si no le asignas el valor a una variable,
	// el programa asigna un valor por defecto (no coloca nulos)

	//SUMA
	x := 10
	y := 50

	result := x + y
	fmt.Println("Suma: ", result)

	//RESTA
	result = y - x
	fmt.Println("Resta:", result)

	//multiplicacion
	result = x * y
	fmt.Println("Multiplicacion:", result)

	//division
	result = y / x
	fmt.Println("Division:", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	//Incremental
	x++
	fmt.Println("Incremental:", x)

	//Decremental
	x--
	fmt.Println("Decremental:", x)

	//Area rectangulo
	baseRectangulo := 20
	alturaRectangulo := 10
	areaRectangulo := baseRectangulo * alturaRectangulo
	fmt.Println("Area rectangulo:", areaRectangulo)

	//Area del circulo
	pi := 3.14
	radioCirculo := 10
	areaCirculo := pi * math.Pow(float64(radioCirculo), 2)
	fmt.Println("Area circulo: ", areaCirculo)

	//Area Trapecio
	baseUno := 6
	baseDos := 15
	alturaTrapecio := 25
	areaTrapecio := ((baseUno + baseDos) * alturaTrapecio) / 2
	fmt.Println("Area Trapecio: ", areaTrapecio)

	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 255
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	//numeros complejos
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i

}
