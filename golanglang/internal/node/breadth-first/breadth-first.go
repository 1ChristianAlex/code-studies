package breadthfirst

import (
	"slices"

	"github.com/1ChristianAlex/code-studies/internal/node"
)

type NodeBreadth struct {
	cumulator []int
}

func (here *NodeBreadth) setCumulator(cumulator int) {
	if !slices.Contains(here.cumulator, cumulator) {
		here.cumulator = append(here.cumulator, cumulator)
	}
}

func (here *NodeBreadth) BreadthFirst(nodeItem node.ListNode[int]) []int {
	queue := make([]node.ListNode[int], 0)

	queue = append(queue, nodeItem)

	for len(queue) > 0 {
		firstItem := queue[0]
		queue = queue[1:]

		here.setCumulator(firstItem.Value)

		if len(firstItem.Next) > 0 {
			for _, value := range firstItem.Next {
				queue = append(queue, *value)
			}
		}
	}

	return here.cumulator
}

func NewNodeBreadth() *NodeBreadth {
	return &NodeBreadth{
		cumulator: make([]int, 0),
	}
}
