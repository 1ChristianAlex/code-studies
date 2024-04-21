package insertmiddle_test

import (
	"reflect"
	"testing"

	insertmiddle "github.com/1ChristianAlex/code-studies/internal/arrays/insert-middle"
)

func TestInsertMiddleArray(t *testing.T) {
	if !reflect.DeepEqual(insertmiddle.InsertMiddleArray([]int{1, 2, 3, 9, 5}, 8), []int{1, 2, 3, 8, 9, 5}) {
		t.Fatalf("Test fail")
	}
	if !reflect.DeepEqual(insertmiddle.InsertMiddleArray([]int{1, 2, 3, 4}, 9), []int{1, 2, 9, 3, 4}) {
		t.Fatalf("Test fail")
	}
}
