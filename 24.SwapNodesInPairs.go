package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummyNode := &ListNode{0, head}
	previous, current := dummyNode, head

	for current != nil && current.Next != nil {
		nextPair := current.Next.Next
		second := current.Next

		second.Next = current
		current.Next = nextPair
		previous.Next = second

		previous = current
		current = nextPair
	}
	return dummyNode.Next
}
