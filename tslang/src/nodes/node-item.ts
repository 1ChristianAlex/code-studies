class Node<T> {
	constructor(
		public value: T,
		public next: Node<T>,
	) {}
}

class ListNode<T> {
	constructor(
		public value: T,
		public next: ListNode<T>[] = [],
	) {}
}

export { Node, ListNode };
