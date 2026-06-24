package main

import "testing"

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		target   int
		expected int
	}{
		{"found first", []int{1, 2, 3, 4, 5}, 1, 0},
		{"found middle", []int{1, 2, 3, 4, 5}, 3, 2},
		{"found last", []int{1, 2, 3, 4, 5}, 5, 4},
		{"not found", []int{1, 2, 3, 4, 5}, 99, -1},
		{"empty array", []int{}, 1, -1},
		{"single element hit", []int{42}, 42, 0},
		{"single element miss", []int{42}, 7, -1},
		{"duplicates returns first", []int{3, 3, 3}, 3, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LinearSearch(tt.array, tt.target)
			if got != tt.expected {
				t.Errorf("LinearSearch(%v, %v) = %d, want %d", tt.array, tt.target, got, tt.expected)
			}
		})
	}
}

func TestLinearSearchStrings(t *testing.T) {
	a := []string{"foo", "bar", "baz"}

	if got := LinearSearch(a, "bar"); got != 1 {
		t.Errorf("got %d, want 1", got)
	}

	if got := LinearSearch(a, "qux"); got != -1 {
		t.Errorf("got %d, want -1", got)
	}
}

func TestLinearSearchRunes(t *testing.T) {
	a := []rune{'z', 'e', 'b', 'c', 'a'}

	if got := LinearSearch(a, 'b'); got != 2 {
		t.Errorf("got %d, want 2", got)
	}

	if got := LinearSearch(a, 'm'); got != -1 {
		t.Errorf("got %d, want -1", got)
	}
}
