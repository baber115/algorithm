package list_node

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(values []int) *ListNode {
	dummy := &ListNode{}
	var tail *ListNode
	tail = dummy
	for _, v := range values {
		node := ListNode{
			Val: v,
		}
		(*tail).Next = &node
		tail = &node
	}

	return dummy.Next
}

func ShowNode(node *ListNode) { //遍历
	for node != nil {
		fmt.Println(*node)
		node = node.Next //移动指针
	}
}
