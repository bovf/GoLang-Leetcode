package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists23(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	head := &ListNode{0, nil}
	newList := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			newList.Next = list1
			list1 = list1.Next
		} else {
			newList.Next = list2
			list2 = list2.Next
		}
		newList = newList.Next
	}
	if list1 != nil {
		newList.Next = list1
	} else {
		newList.Next = list2
	}

	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	newListNotNull := []*ListNode{}
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		newListNotNull = append(newListNotNull, lists[i])
	}

	if len(newListNotNull) == 0 {
		return nil
	}
	mergedList := newListNotNull[0]

	for i := 1; i < len(newListNotNull); i++ {
		if newListNotNull[i] != nil {
			mergedList = mergeTwoLists23(mergedList, newListNotNull[i])
		}
	}
	return mergedList
}
