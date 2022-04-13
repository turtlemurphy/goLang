package main

import (
	"fmt"
	"sort"
)

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

	//In fact, the Go spec says that the iteration order over maps is not specified.
	//That is to say, you should not expect the map keys to appear in any particular
	//order. This makes sense when you understand how Go maps are implemented, but we
	//needn't worry about the details here. Just know that you can't rely on iterating over a map in a specific order.

	kvs := map[string]string{"a": "apple", "b": "banana", "c": "cherry", "e": "egg", "f": "fig"}

	//Uses Println to print out all the key/element pair in order
	fmt.Println("First print using Println")
	fmt.Println(kvs)
	fmt.Print("\n")

	//This loop will print all the key/element pairs but not always in the same order due to the nature of the map data structure
	fmt.Println("Second Printing using a for loop, will not always be in alphabetical order")
	for k, element := range kvs {
		//fmt.Println("key:", key)
		fmt.Printf("%s -> %s\n", k, element)
	}
	fmt.Print("\n")

	//This loop will take the keys, sort them into an array by alphbetical order then print out the map elements in oprder from the sorted array
	var keys []string
	for key := range kvs {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	fmt.Println(keys)
	fmt.Println("In alphabetical order:")
	for _, key := range keys {
		fmt.Printf("%s -> %s\n", key, kvs[key])
		//fmt.Println(key, kvs[key])
	}

}
