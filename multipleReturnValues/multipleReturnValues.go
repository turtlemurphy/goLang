package main

import "fmt"

func vals() (int, int, int, int, int, int) {
	return 3, 7, 2, 6, 9, 4
}
func main() {

	a, b, c, d, e, f := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	// _, j := vals()
	// fmt.Println(j)
}
