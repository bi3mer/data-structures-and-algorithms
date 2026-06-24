package main

import "fmt"

func LinearSearch[T comparable](array []T, target T) int {
	for i, v := range array {
		if v == target {
			return i
		}
	}

	return -1
}

func main() {
	a := []rune{'z', 'e', 'b', 'c', 'a'}
	target := rune('m')
	res := LinearSearch(a, target)
	if res == -1 {
		fmt.Println("target not found!")
	} else {
		fmt.Printf("Target found at index %d\n", res)
	}
}
