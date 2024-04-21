package ktlang.array.insert.middle

fun <T> arrayInsertMiddle(list: List<T>, newValue: T): List<T> {
    val newList: MutableList<T> = mutableListOf<T>()

    val middleIndex = Math.ceil(list.size.toDouble() / 2).toInt()

    for (i in list.indices) {
        if (i == middleIndex) {
            newList.add(newValue)
            newList.add(list[i])
        } else {
            newList.add(list[i])
        }
    }

    return newList
}
