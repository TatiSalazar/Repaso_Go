/**
package main

import "fmt"

// Un channel es un tipo de variable que permite el paso de información
// entre gorourines.

func say(text string, c chan<- string) { // <- ch De esta forma extraemos datos del channel.
	c <- text // introducimos informacion al channel.
}

func main() {
	c := make(chan string, 1) //declaración del channel
	fmt.Println("Hello")
	go say("Bye", c)
	fmt.Println(<-c)
}
