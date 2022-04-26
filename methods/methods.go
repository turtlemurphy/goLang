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
		if b[i] {
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

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b uint) uint {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	fmt.Println(a)
	return a
}

// Euler’s Totient Function
func Phi(n uint) (result uint) {
	// Initialize result as n
	result = n

	// Consider all prime factors of n and subtract their
	// multiples from result
	var p uint
	for p = 2; p*p <= n; p++ {
		// Check if p is a prime factor.
		if n%p == 0 {
			// If yes, then update n and result
			for n%p == 0 {
				n /= p
			}
			result -= (result / p)
		}
	}

	// If n has a prime factor greater than sqrt(n)
	// (There can be at-most one such prime factor)
	if n > 1 {
		result -= (result / n)
	}

	fmt.Println(result)

	return result
}

func main() {

	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	fmt.Println()

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	fmt.Println()

	sieveOfEratosthenes(42069)

	fmt.Println()

	GCD(7231596374929284751, 69427)

	fmt.Println()

	Phi(8)

}
