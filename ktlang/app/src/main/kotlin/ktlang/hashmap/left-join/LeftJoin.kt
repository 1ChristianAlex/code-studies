package ktlang.hashmap.left.join

fun leftJoin(leftMap: Map<String, String>, rightMap: Map<String, String>): List<List<String?>> {

    val newList = mutableListOf<List<String?>>()

    for (entry in leftMap.entries) {
        rightMap.get(entry.key).let {
            newList.add(
                    listOf<String?>(entry.key, entry.value, if (it.isNullOrEmpty()) null else it)
            )
        }
    }

    return newList
}
