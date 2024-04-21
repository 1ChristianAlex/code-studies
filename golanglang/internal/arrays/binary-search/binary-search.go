package binarysearch

func ArrayBinarySearch(list []int, target int) int {
	lowestIndex := 0
	highestIndex := len(list) - 1

	for lowestIndex <= highestIndex {
		middleIndex := (lowestIndex + highestIndex) / 2
		middleItem := list[middleIndex]

		if middleItem == target {
			return middleIndex
		}

		if middleItem > target {
			highestIndex = middleIndex - 1
		} else {
			lowestIndex = middleIndex + 1
		}
	}

	return -1
}
