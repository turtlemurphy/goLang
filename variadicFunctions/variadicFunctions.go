package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Print("\n")
	fmt.Println(total)
}
func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1}

	for j := 1; j <= 420; j++ {
		sum(nums...)
		nums = append(nums, j)
	}
}
