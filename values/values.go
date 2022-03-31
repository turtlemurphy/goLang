package main

import "fmt"

func main() {

	val := 122

	fmt.Printf("%d\n", val)
	fmt.Printf("%c\n", val)
	fmt.Printf("%q\n", val)
	fmt.Printf("%x\n", val)
	fmt.Printf("%X\n", val)
	fmt.Printf("%o\n", val)
	fmt.Printf("%O\n", val)
	fmt.Printf("%b\n", val)
	fmt.Printf("%U\n", val)

	fmt.Printf("%d\n", 1671)
	fmt.Printf("%o\n", 1671)
	fmt.Printf("%x\n", 1671)
	fmt.Printf("%X\n", 1671)
	fmt.Printf("%#b\n", 1671)
	fmt.Printf("%f\n", 1671.678)
	fmt.Printf("%F\n", 1671.678)
	fmt.Printf("%e\n", 1671.678)
	fmt.Printf("%E\n", 1671.678)
	fmt.Printf("%g\n", 1671.678)
	fmt.Printf("%G\n", 1671.678)
	fmt.Printf("%s\n", "Zetcode")
	fmt.Printf("%c %c %c %c %c %c %c\n", 'Z', 'e', 't', 'C', 'o', 'd', 'e')
	fmt.Printf("%p\n", []int{1, 2, 3})
	fmt.Printf("%d%%\n", 1671)
	fmt.Printf("%t\n", 3 > 5)
	fmt.Printf("%t\n", 5 > 3)

	fmt.Println("go" + "lang")

	fmt.Println("6 + 9 + 4 + 2 + 0 =", 6+9+4+2+0)
	fmt.Println("255.3/3.7 =", 255.3/3.7)

	fmt.Print("true AND false OR false AND true\n", (true && false) || (false && true), "\n")
	fmt.Print("true OR false AND false OR true\n", (true || false) && (false || true), "\n")
	fmt.Print("true OR false OR false OR false OR false OR false OR false OR false\n", (true || false || false || false || false || false || false || false), "\n")
	fmt.Print("NOT false AND false AND false AND false AND false AND false\n", !(false && false && false && false && false && false), "\n")
}
