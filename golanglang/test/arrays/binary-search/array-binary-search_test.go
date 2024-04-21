package binarysearch_test

import (
	"testing"

	binarysearch "github.com/1ChristianAlex/code-studies/internal/arrays/binary-search"
)

func TestArrayBinarySearch(t *testing.T) {
	if binarysearch.ArrayBinarySearch([]int{1, 2, 3, 4, 5, 9, 12}, 8) != -1 {
		t.Fatalf("Test fail")
	}

	if binarysearch.ArrayBinarySearch([]int{1, 2, 3, 4, 5}, 5) != 4 {
		t.Fatalf("Test fail")
	}

	if binarysearch.ArrayBinarySearch([]int{1, 2, 3, 5, 9, 12, 32, 41}, 41) != 7 {
		t.Fatalf("Test fail")
	}
	if binarysearch.ArrayBinarySearch([]int{1, 2, 3, 12, 32, 5}, 1) != 0 {
		t.Fatalf("Test fail")
	}
}
