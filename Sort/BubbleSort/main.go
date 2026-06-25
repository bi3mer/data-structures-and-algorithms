package main

import (
	"cmp"
	"fmt"
)

func BubbleSort[T cmp.Ordered](s []T) {
	n := len(s)

	for i := range n - 1 {
		for j := range n - 1 - i {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func main() {
	a := []rune{'z', 'e', 'b', 'c', 'a', 'd', 'r', 's', 'w'}
	fmt.Println(string(a))
	BubbleSort(a)
	fmt.Println(string(a))
}
