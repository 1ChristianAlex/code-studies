package leftjoin

func DoLeftJoin(hashMapOne map[string]string, hashMapSecond map[string]string) [][]string {
	asSlice := make([][]string, len(hashMapOne))

	index := 0

	for key, value := range hashMapOne {

		asSlice[index] = append(asSlice[index], key, value, hashMapSecond[key])
		index++
	}

	return asSlice
}
