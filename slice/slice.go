package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1 [2:5]:", l)

	l = s[2:3]
	fmt.Println("sl2 [2:3]:", l)

	l = s[2:4]
	fmt.Println("sl3 [2:4]:", l)

	l = s[2:5]
	fmt.Println("sl4 [2:5]:", l)

	l = s[5:5]
	fmt.Println("sl5 [5:5]:", l)

	l = s[:5]
	fmt.Println("sl6 [ :5]:", l)

	l = s[2:]
	fmt.Println("sl7 [2: ]:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl T:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
