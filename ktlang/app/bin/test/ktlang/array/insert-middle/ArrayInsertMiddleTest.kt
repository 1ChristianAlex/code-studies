import kotlin.test.Test
import kotlin.test.assertEquals
import ktlang.array.insert.middle.*

class ArrayInsertMiddleTest {

    @Test
    fun testAcertMiddle() {
        val result = arrayInsertMiddle(listOf(1, 2, 3, 9, 5), 8)
        assertEquals(listOf(1, 2, 3, 8, 9, 5), result)
        // assertEquals(listOf(1, 2, 9, 3, 4), arrayInsertMiddle(listOf(1, 2, 3, 4), 9))
    }
}
