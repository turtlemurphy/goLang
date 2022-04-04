package main

import "fmt"

func main() {

	fmt.Print("\n")
	fmt.Print("\n")

	for i := 1; i <= 69; i++ {

		if i%2 == 0 {
			fmt.Println(i, " is even")
			//fmt.Print("\n")
		} else {
			fmt.Println(i, " is odd")
			//fmt.Print("\n")
		}

		if i%4 == 0 {
			fmt.Println(i, " is divisible by 4")
			//fmt.Print("\n")
		}

		if i < 0 {
			fmt.Println(i, "is negative")
			fmt.Print("\n")
		} else if i < 10 {
			fmt.Println(i, "has 1 digit")
			fmt.Print("\n")
		} else {
			fmt.Println(i, "has multiple digits")
			fmt.Print("\n")
		}
		// if i%20 == 0 {
		// 	fmt.Print("\n")
		// }
	}

}
