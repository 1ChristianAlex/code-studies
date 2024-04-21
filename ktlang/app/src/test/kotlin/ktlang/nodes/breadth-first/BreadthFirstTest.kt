import kotlin.test.Test
import kotlin.test.assertEquals
import ktlang.nodes.breadthFirst.NodeBreadth
import ktlang.nodes.nodeitem.ListNode

class NodeBreadthTest {

    @Test
    fun testBreachFirst() {

        val five = ListNode(5, mutableListOf())
        val four = ListNode(4, mutableListOf())
        val one = ListNode(1, mutableListOf())
        val two = ListNode(2, mutableListOf())
        val three = ListNode(3, mutableListOf())
        val zero = ListNode(0, mutableListOf())

        four.next.add(five)

        one.next.add(four)
        two.next.add(four)
        three.next.add(four)

        zero.next.addAll(listOf(one, two, three))

        val instante = NodeBreadth()
        val result = instante.breadthFirst(zero)

        assertEquals(listOf(0, 1, 2, 3, 4, 5), result)
    }
    @Test
    fun testBreachFirstSecond() {

        val five = ListNode(5, mutableListOf())
        val four = ListNode(4, mutableListOf())
        val six = ListNode(6, mutableListOf())
        val two = ListNode(2, mutableListOf())
        val three = ListNode(3, mutableListOf())
        val zero = ListNode(0, mutableListOf())

        four.next.add(five)

        two.next.add(four)

        zero.next.addAll(listOf(six, two, three))

        val instante = NodeBreadth()
        val result = instante.breadthFirst(zero)

        assertEquals(listOf(0, 6, 2, 3, 4, 5), result)
    }

    @Test
    fun testBreachFirstThrid() {

        val zero = ListNode(0, mutableListOf())
        val one = ListNode(1, mutableListOf())
        val two = ListNode(2, mutableListOf())
        val three = ListNode(3, mutableListOf())
        val four = ListNode(4, mutableListOf())
        val five = ListNode(5, mutableListOf())
        val six = ListNode(6, mutableListOf())
        val seven = ListNode(7, mutableListOf())
        val eight = ListNode(8, mutableListOf())

        zero.next.addAll(listOf(one, two, three))
        one.next.add(four)
        two.next.add(four)
        three.next.add(four)

        four.next.addAll(listOf(seven, five))
        five.next.addAll(listOf(seven, six))
        seven.next.add(eight)

        val instante = NodeBreadth()
        val result = instante.breadthFirst(zero)

        assertEquals(listOf(0, 1, 2, 3, 4, 7, 5, 8, 6), result)
    }
}
