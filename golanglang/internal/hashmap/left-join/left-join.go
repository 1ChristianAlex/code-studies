package leftjoin

import "sort"

func DoLeftJoin(hashMapOne map[string]string, hashMapSecond map[string]string) [][]string {
	asSlice := make([][]string, len(hashMapOne))
	sliceKeys := make([]string, 0)

	for key := range hashMapOne {
		sliceKeys = append(sliceKeys, key)
	}

	sort.Strings(sliceKeys)

	index := 0
	for _, key := range sliceKeys {

		asSlice[index] = append(asSlice[index], key, hashMapOne[key], hashMapSecond[key])
		index++
	}

	return asSlice
}
