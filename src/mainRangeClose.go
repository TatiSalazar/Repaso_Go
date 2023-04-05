package main

import "fmt"

func mesage(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "mensaje1"
	c <- "mensaje2"

	fmt.Println(len(c), cap(c))

	// range y close
	close(c)
	//c <- "mensaje3"

	for message := range c {
		fmt.Println(message)
	}

	// select
	email1 := make(chan string)
	email2 := make(chan string)

	go mesage("mensaje1", email1)
	go mesage("mensaje2", email2)

	for i:=0 , i < 2; i ++{
		select {
		case ml := <-email:
			fmt.
		}

	}
}
