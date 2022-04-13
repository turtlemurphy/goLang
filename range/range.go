package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	sum := 0

	for _, num := range nums {
		sum += num
		fmt.Println("sum:", sum)
	}

	fmt.Print("\n")

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	fmt.Print("\n")

	kvs := map[string]string{"a": "apple", "b": "banana", "c": "cherry"}
	//kvs := map[string]string{"a": "apple", "b": "banana", "c": "cherry", "d": "date", "e": "egg", "f": "fig"}

	for k, v := range kvs {
		fmt.Println("key:", k)
		fmt.Printf("%s -> %s\n", k, v)
	}
}
