package node

type ListNode[T any] struct {
	Value T
	Next  []*ListNode[T]
}
