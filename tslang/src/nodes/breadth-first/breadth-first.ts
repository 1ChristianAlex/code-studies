import type { ListNode, Node } from "../node-item";

class NodeBreadth {
	private readonly cumulator = new Set<number>();

	breadthFirst = (nodeItem: ListNode<number>): number[] => {
		const currentNode = nodeItem;

		const queue = [currentNode];

		while (queue.length) {
			const currentQueue = queue.shift();

			if (currentQueue) {
				this.cumulator.add(currentQueue.value);

				queue.push(...currentQueue.next);
			}
		}

		const result = Array.from(this.cumulator);

		return result;
	};
}

export { NodeBreadth };
