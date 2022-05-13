package main

import "fmt"

func main() {

	messages := make(chan string, 9)

	//Makes a FIFO stack
	messages <- "the"
	messages <- "quick"
	messages <- "brown"
	messages <- "fox"
	messages <- "jumped"
	messages <- "over"
	messages <- "the"
	messages <- "lazy"
	messages <- "dog"

	//Pops the stack
	//failure if you pop too many messages off the stack

	for i := 1; i <= 9; i++ {

		fmt.Println(<-messages)

	}

}
