package main

import "fmt"

func main() {

	var T [50]int
	fmt.Println("T[] is declared and empty: ", T)
	fmt.Print("\n")

	fmt.Println("Set T[0], aka slot 1, to 69")
	fmt.Println("Set T[41], aka slot 42, to 100")
	fmt.Print("\n")

	T[0] = 69
	T[41] = 100

	fmt.Println("Set:", T)
	fmt.Print("\n")

	fmt.Println("Get:", T[41])

	fmt.Println("len:", len(T))
	fmt.Print("\n")

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b is declared and its slots have been pre initialized", b)
	fmt.Print("\n")

	var twoD [20][20]int
	for i := 0; i < 19; i++ {
		for j := 0; j < 19; j++ {
			twoD[i][j] = i + j
			fmt.Print(twoD[i][j], " ")
		}
		fmt.Print("\n")
	}
	//fmt.Println("twoD: ", twoD)
}
