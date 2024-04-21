package ktlang.array.binary.search

fun arrayBinarySearch(list: List<Int>, target: Int): Int {
    var lowestIndex = 0
    var highestIndex = list.size - 1

    while (lowestIndex <= highestIndex) {
        val middleIndex = ((highestIndex + lowestIndex) / 2).toInt()

        val currentMiddle = list[middleIndex]

        if (currentMiddle == target) {
            return middleIndex
        }

        if (currentMiddle > target) {
            highestIndex = middleIndex - 1
        } else {
            lowestIndex = middleIndex + 1
        }
    }

    return -1
}
