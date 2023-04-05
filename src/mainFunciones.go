/**
package main

import "fmt"

func normalFuntion(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func areaRectangulo(base, altura float64) float64 {
	return base * altura
}

func areaCirculo(radio float64) float64 {
	const pi float64 = 3.14
	return pi * (radio * radio)
}

func areaTrapecio(a, b, h float64) float64 {
	return (((a + b) * h) / 2)
}

func main() {
	normalFuntion("Hola Mundo")
	tripleArgument(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("Value: ", value)

	value1, value2 := doubleReturn(2)
	//value1, _ := doubleReturn(2)
	fmt.Println("Value1 : ", value1, value2)

	// rectangulo
	rectangulo := areaRectangulo(4, 10)
	//circulo
	circulo := areaCirculo(3.455)
	//trapecio
	trapecio := areaTrapecio(14, 20, 5)

	message := fmt.Sprintf("El area del rectangulo es: %f \nEl area del circulo: %f \nEl area del trapecio %f", rectangulo, circulo, trapecio)
	fmt.Println(message)
}
