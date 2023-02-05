package list_node

import "testing"

// https://leetcode.cn/problems/swap-nodes-in-pairs/
func TestSwapPairs(t *testing.T) {

}

func swapPairs(head *ListNode) *ListNode {
	dummyListNode := &ListNode{Next: head}
	pre := dummyListNode
	for pre.Next != nil && pre.Next.Next != nil {
		node1 := pre.Next
		node2 := pre.Next.Next
		pre.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		pre = node1
	}

	return dummyListNode.Next
}

func swapPairsRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairsRecursion(newHead.Next)
	newHead.Next = head

	return newHead
}
