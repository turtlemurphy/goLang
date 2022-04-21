package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()
	newInts := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Print("\n")
	fmt.Print("\n")

	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts())
}
