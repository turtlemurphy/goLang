package main

import "fmt"

func main() {

	val := 122
	//The integer is formatted in decimal, hexadecimal, octal,
	//and binary notations. It is also formatted as a character
	// literal, quoted charracter literal, and in a Unicode escape sequence.
	fmt.Printf("%d\n", val)
	fmt.Printf("%c\n", val)
	fmt.Printf("%q\n", val)
	fmt.Printf("%x\n", val)
	fmt.Printf("%X\n", val)
	fmt.Printf("%o\n", val)
	fmt.Printf("%O\n", val)
	fmt.Printf("%b\n", val)
	fmt.Printf("%U\n", val)

	fmt.Printf("%0.f\n", 16.540)
	fmt.Printf("%0.2f\n", 16.540)
	fmt.Printf("%0.3f\n", 16.540)
	fmt.Printf("%0.5f\n", 16.540)

	//f - decimal floating point, lowercase
	//F - decimal floating point, uppercase
	//e - scientific notation (mantissa/exponent), lowercase
	//E - scientific notation (mantissa/exponent), uppercase
	//g - uses the shortest representation of %e or %f
	//G - Use the shortest representation of %E or %F

	val1 := 1273.78888769000

	fmt.Printf("%f\n", val1)
	fmt.Printf("%e\n", val1)
	fmt.Printf("%g\n", val1)
	fmt.Printf("%E\n", val1)
	fmt.Printf("%G\n", val1)

	fmt.Println("-------------------------")

	fmt.Printf("%.10f\n", val1)
	fmt.Printf("%.10e\n", val1)
	fmt.Printf("%.10g\n", val1)
	fmt.Printf("%.10E\n", val1)
	fmt.Printf("%.10G\n", val1)

	fmt.Println("-------------------------")

	val3 := 66_000_000_000.1200

	fmt.Printf("%f\n", val3)
	fmt.Printf("%e\n", val3)
	fmt.Printf("%g\n", val3)
	fmt.Printf("%E\n", val3)
	fmt.Printf("%G\n", val3)

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
