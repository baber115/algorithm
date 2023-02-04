package list_node

import (
	"testing"
)

// https://leetcode.cn/problems/remove-linked-list-elements/
func TestRemoveLinkedListElement(t *testing.T) {
}

// 递归法
func removeElementsRecursion(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElementsRecursion(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}

// 基本操作
func removeElementsNormal(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	for cur := dummyHead; cur.Next != nil; {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummyHead.Next
}
