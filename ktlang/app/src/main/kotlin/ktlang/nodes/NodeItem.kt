package ktlang.nodes.nodeitem

class ListNode<T>(public val value: T, public val next: MutableList<ListNode<T>>) {}
