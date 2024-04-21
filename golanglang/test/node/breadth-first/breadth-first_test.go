package breadthfirst

import (
	"reflect"
	"testing"

	node "github.com/1ChristianAlex/code-studies/internal/node"
	breadthfirst "github.com/1ChristianAlex/code-studies/internal/node/breadth-first"
)

func TestBreadthFirst(t *testing.T) {
	five := &node.ListNode[int]{Value: 5}
	four := &node.ListNode[int]{Value: 4}
	one := &node.ListNode[int]{Value: 1}
	two := &node.ListNode[int]{Value: 2}
	three := &node.ListNode[int]{Value: 3}
	zero := &node.ListNode[int]{Value: 0}

	four.Next = append(four.Next, five)

	one.Next = append(one.Next, four)
	two.Next = append(two.Next, four)
	three.Next = append(three.Next, four)

	zero.Next = append(zero.Next, one, two, three)

	if !reflect.DeepEqual(breadthfirst.NewNodeBreadth().BreadthFirst(*zero), []int{0, 1, 2, 3, 4, 5}) {
		t.Fatalf("Test fail")
	}
}

func TestBreadthFirstSecond(t *testing.T) {
	six := &node.ListNode[int]{Value: 6}
	five := &node.ListNode[int]{Value: 5}
	four := &node.ListNode[int]{Value: 4}
	two := &node.ListNode[int]{Value: 2}
	three := &node.ListNode[int]{Value: 3}
	zero := &node.ListNode[int]{Value: 0}

	four.Next = append(four.Next, five)

	two.Next = append(two.Next, four)

	zero.Next = append(zero.Next, six, two, three)

	if !reflect.DeepEqual(breadthfirst.NewNodeBreadth().BreadthFirst(*zero), []int{0, 6, 2, 3, 4, 5}) {
		t.Fatalf("Test fail")
	}
}

func TestBreadthFirstThrid(t *testing.T) {
	zero := &node.ListNode[int]{Value: 0}
	one := &node.ListNode[int]{Value: 1}
	two := &node.ListNode[int]{Value: 2}
	three := &node.ListNode[int]{Value: 3}
	four := &node.ListNode[int]{Value: 4}
	five := &node.ListNode[int]{Value: 5}
	six := &node.ListNode[int]{Value: 6}
	seven := &node.ListNode[int]{Value: 7}
	eight := &node.ListNode[int]{Value: 8}

	zero.Next = append(zero.Next, one, two, three)
	one.Next = append(one.Next, four)
	two.Next = append(two.Next, four)
	three.Next = append(three.Next, four)

	four.Next = append(four.Next, seven, five)
	five.Next = append(five.Next, seven, six)
	seven.Next = append(seven.Next, eight)

	if !reflect.DeepEqual(breadthfirst.NewNodeBreadth().BreadthFirst(*zero), []int{0, 1, 2, 3, 4, 7, 5, 8, 6}) {
		t.Fatalf("Test fail")
	}
}
