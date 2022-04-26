package main

import "fmt"

type rect struct {
	width, height int
}

//func 1 - AREA
func (r *rect) area() int {
	return r.width * r.height
}

//func 2 - Referenced AREA
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

//func 3 - Prime Number Generator
func sieveOfEratosthenes(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	fmt.Println(primes)
	return
}

func main() {

	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	sieveOfEratosthenes(10000)

}
