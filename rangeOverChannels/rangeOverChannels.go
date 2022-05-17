package main

import "fmt"

func main() {

	queue := make(chan string, 8)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	queue <- "four"
	queue <- "five"
	queue <- "six"
	queue <- "seven"
	queue <- "eight"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
