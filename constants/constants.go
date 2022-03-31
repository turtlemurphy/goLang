package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const b = 42
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int8(b))
	fmt.Println(int16(b))
	fmt.Println(int32(b))
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
	fmt.Println(math.Cos(n))
	fmt.Println(math.Tan(n))
}
