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
		{"diligent", "employed", "idle"},
		{"font", "enamored", "averse"},
		{"guide", "usher", "follow"},
		{"outfit", "garb", ""},
		{"wrath", "anger", "delight"},
	}

	joinResult := leftjoin.DoLeftJoin(hashMapOne, hashMapTwo)

	if !reflect.DeepEqual(joinResult, hashMapResult) {
		t.Fatalf("Test fail")
	}
}
