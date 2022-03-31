package main

import "fmt"

func main() {

	i := 1
	for i <= 420 {

		fmt.Print(i, " ")
		i = i + 1
		if i%20 == 0 {
			fmt.Print("\n")
		}
	}

	fmt.Print("\n")

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
