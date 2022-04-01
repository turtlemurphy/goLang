package main

import "fmt"

// return list of primes less than N
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
	return
}

func main() {

	primes := sieveOfEratosthenes(1000)
	ind := 1
	//underscore operator = blank identifier
	for _, p := range primes {
		fmt.Print(p, " ")

		if ind%20 == 0 {
			fmt.Print("\n")
		}
		ind++
	}

	fmt.Print("\n")
	fmt.Print("\n")

	for i := 1; i <= 42; i++ {
		fmt.Print(i, " ")
		i++
		if i%20 == 0 {
			fmt.Print("\n")
		}
	}

	fmt.Print("\n")
	fmt.Print("\n")

	for j := 0; j <= 10; j++ {
		fmt.Println(j)
	}

	fmt.Print("\n")
	fmt.Print("\n")

	for {
		fmt.Println("loop")
		break
	}

	fmt.Print("\n")
	fmt.Print("\n")

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	fmt.Print("\n")
	fmt.Print("\n")
}
