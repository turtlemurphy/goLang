package main

import "fmt"

func main() {

	fmt.Print("\n")
	fmt.Print("\n")

	for i := -100; i <= 100; i += 10 {

		if i%2 == 0 {
			fmt.Println(i, " is even")

		} else {
			fmt.Println(i, " is odd")

		}

		if i%4 == 0 {
			fmt.Println(i, " is divisible by 4")

		}

		if i < 0 {
			fmt.Println(i, " is negative")

		} else if i < 10 {
			fmt.Println(i, " has 1 digit")

		} else {
			fmt.Println(i, " has multiple digits")

		}
		// if i%20 == 0 {
		// 	fmt.Print("\n")
		// }
	}

}
