package main

import "fmt"

func main() {

	messages := make(chan string, 9)

	messages <- "the"
	messages <- "quick"
	messages <- "brown"
	messages <- "fox"
	messages <- "jumped"
	messages <- "over"
	messages <- "the"
	messages <- "lazy"
	messages <- "dog"

	for i := 1; i < 9; i++ {

		fmt.Println(<-messages)

	}

}
