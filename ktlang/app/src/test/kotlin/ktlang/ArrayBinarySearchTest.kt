import kotlin.test.Test
import kotlin.test.assertEquals
import ktlang.array.binary.search.arrayBinarySearch

class ArrayBinarySearchTest {

    @Test
    fun findTarget() {
        assertEquals(-1, arrayBinarySearch(listOf(1, 2, 3, 4, 5, 9, 12), 8))
        assertEquals(4, arrayBinarySearch(listOf(1, 2, 3, 4, 5), 5))
        assertEquals(7, arrayBinarySearch(listOf(1, 2, 3, 5, 9, 12, 32, 41), 41))
        assertEquals(0, arrayBinarySearch(listOf(1, 2, 3, 12, 32, 5), 1))
    }
}
