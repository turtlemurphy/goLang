package main

import "fmt"

func main() {

	fmt.Print("\n")
	fmt.Print("\n")

	if 7%2 == 0 {
		fmt.Println("7 is even")
		fmt.Print("\n")
	} else {
		fmt.Println("7 is odd")
		fmt.Print("\n")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
		fmt.Print("\n")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
		fmt.Print("\n")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
		fmt.Print("\n")
	} else {
		fmt.Println(num, "has multiple digits")
		fmt.Print("\n")
	}
}
