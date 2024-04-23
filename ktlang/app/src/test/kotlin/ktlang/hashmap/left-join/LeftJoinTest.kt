import kotlin.test.Test
import kotlin.test.assertEquals
import ktlang.hashmap.left.join.leftJoin

class LeftJoinTest {

    @Test
    fun ShouldLeftJoin() {

        val hashMapOne =
                mapOf(
                        "diligent" to "employed",
                        "font" to "enamored",
                        "guide" to "usher",
                        "outfit" to "garb",
                        "wrath" to "anger",
                )

        val hashMapTwo =
                mapOf(
                        "diligent" to "idle",
                        "font" to "averse",
                        "guide" to "follow",
                        "flow" to "jam",
                        "wrath" to "delight",
                )

        val hashMapExpected =
                listOf(
                        listOf("diligent", "employed", "idle"),
                        listOf("font", "enamored", "averse"),
                        listOf("guide", "usher", "follow"),
                        listOf("outfit", "garb", null),
                        listOf("wrath", "anger", "delight"),
                )

        val mapResult = leftJoin(hashMapOne, hashMapTwo)

        assertEquals(mapResult, hashMapExpected)
    }
}
