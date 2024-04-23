package leftjoin_test

import (
	"reflect"
	"testing"

	leftjoin "github.com/1ChristianAlex/code-studies/internal/hashmap/left-join"
)

func TestDoLeftJoin(t *testing.T) {
	hashMapOne := map[string]string{
		"diligent": "employed",
		"font":     "enamored",
		"guide":    "usher",
		"outfit":   "garb",
		"wrath":    "anger",
	}

	hashMapTwo := map[string]string{
		"diligent": "idle",
		"font":     "averse",
		"guide":    "follow",
		"flow":     "jam",
		"wrath":    "delight",
	}

	hashMapResult := [][]string{
		{"wrath", "anger", "delight"},
		{"diligent", "employed", "idle"},
		{"font", "enamored", "averse"},
		{"guide", "usher", "follow"},
		{"outfit", "garb", ""},
	}

	if !reflect.DeepEqual(leftjoin.DoLeftJoin(hashMapOne, hashMapTwo), hashMapResult) {
		t.Fatalf("Test fail")
	}
}
