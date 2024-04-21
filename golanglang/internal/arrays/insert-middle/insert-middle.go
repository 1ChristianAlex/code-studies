package insertmiddle

import "math"

func InsertMiddleArray[T interface{}](list []T, newValue T) []T {
	newList := make([]T, len(list)+1)

	middleIndex := int(math.Ceil(float64(len(list)) / 2))

	newList[middleIndex] = newValue

	for indexList, value := range list {
		if indexList >= middleIndex {
			newList[indexList+1] = value
		} else {
			newList[indexList] = value
		}
	}

	return newList
}
