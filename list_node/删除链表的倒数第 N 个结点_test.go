package list_node

import "testing"

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
func TestRemoveNthFromEnd(t *testing.T) {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slowListNode, fastListNode := dummy, dummy

	for i := 0; i < n+1; i++ {
		fastListNode = fastListNode.Next
	}
	for fastListNode != nil {
		fastListNode = fastListNode.Next
		slowListNode = slowListNode.Next
	}
	slowListNode.Next = slowListNode.Next.Next

	return dummy.Next
}

func removeNthFromEndII(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slowListNode, fastListNode := dummy, head

	i := 1
	for fastListNode != nil {
		fastListNode = fastListNode.Next
		if i > n {
			slowListNode = slowListNode.Next
		}
		i++
	}
	slowListNode.Next = slowListNode.Next.Next

	return slowListNode
}
