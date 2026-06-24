package main

import (
	"cmp"
	"fmt"
)

func SelectionSort[T cmp.Ordered](s []T) {
	n := len(s)

	for i := range n - 1 {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if s[j] < s[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			s[i], s[minIndex] = s[minIndex], s[i]
		}
	}
}

func main() {
	a := []rune{'z', 'e', 'b', 'c', 'a', 'd', 'r', 's', 'w'}
	fmt.Println(string(a))
	SelectionSort(a)
	fmt.Println(string(a))
}
