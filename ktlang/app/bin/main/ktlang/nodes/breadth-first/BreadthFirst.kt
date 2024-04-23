package ktlang.nodes.breadthFirst

import ktlang.nodes.nodeitem.ListNode

class NodeBreadth {
    private val cumulator = mutableSetOf<Int>()

    fun breadthFirst(nodes: ListNode<Int>): List<Int> {

        val queue = mutableListOf(nodes)

        while (queue.size > 0) {
            val currentItem = queue.removeAt(0)
            cumulator.add(currentItem.value)

            if (currentItem.next.size > 0) {
                queue.addAll(currentItem.next)
            }
        }

        return cumulator.toList()
    }
}
