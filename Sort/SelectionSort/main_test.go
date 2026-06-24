package main

import (
	"slices"
	"testing"
)

func TestSelectionSortInts(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse sorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"unsorted", []int{3, 1, 4, 1, 5, 9, 2, 6}, []int{1, 1, 2, 3, 4, 5, 6, 9}},
		{"single element", []int{42}, []int{42}},
		{"empty", []int{}, []int{}},
		{"duplicates", []int{3, 3, 3}, []int{3, 3, 3}},
		{"two elements sorted", []int{1, 2}, []int{1, 2}},
		{"two elements unsorted", []int{2, 1}, []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Clone(tt.input)
			SelectionSort(got)
			if !slices.Equal(got, tt.expected) {
				t.Errorf("SelectionSort(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestSelectionSortStrings(t *testing.T) {
	a := []string{"banana", "apple", "cherry"}
	SelectionSort(a)
	expected := []string{"apple", "banana", "cherry"}

	if !slices.Equal(a, expected) {
		t.Errorf("got %v, want %v", a, expected)
	}
}

func TestSelectionSortRunes(t *testing.T) {
	a := []rune{'z', 'e', 'b', 'c', 'a'}
	SelectionSort(a)
	expected := []rune{'a', 'b', 'c', 'e', 'z'}

	if !slices.Equal(a, expected) {
		t.Errorf("got %s, want %s", string(a), string(expected))
	}
}
