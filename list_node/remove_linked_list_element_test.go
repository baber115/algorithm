package list_node

import (
	"testing"
)

func TestRemoveLinkedListElement(t *testing.T) {
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}

//func removeElements(head *ListNode, val int) *ListNode {
//	dummyHead := &ListNode{Next: head}
//	for cur := dummyHead; cur.Next != nil; {
//		if cur.Next.Val == val {
//			cur.Next = cur.Next.Next
//		} else {
//			cur = cur.Next
//		}
//	}
//
//	return dummyHead.Next
//}
