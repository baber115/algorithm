package list_node

import "testing"

// https://leetcode.cn/problems/reverse-linked-list/
func TestReverseLinkedList(t *testing.T) {
}

// 双指针
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	return pre
}

// 方便理解的递归，根据双指针映射
func reverseListRecursionI(head *ListNode) *ListNode {
	return reverse(head, nil)
}

func reverse(cur, pre *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	temp := cur.Next
	cur.Next = pre
	return reverse(temp, cur)
}

// 官方题解递归
// 以链表1->2->3->4->5->nil为例
func reverseListRecursionII(head *ListNode) *ListNode {
	/*
		直到当前节点的下一个节点为空时返回当前节点
		5下一个节点是nil，返回5
	*/
	if head == nil || head.Next == nil {
		return head
	}
	// 递归传入下一个节点，目的是为了到达最后一个节点
	newHead := reverseListRecursionII(head.Next)
	/*
		第一轮出栈：head=5, head.Next=nil,return 5
		第二轮出栈：head=4,head.Next=5
				执行head.Next.Next=head,5.Next=4,把5的下一个节点指向4。此时链表为1->2->3->4<->5
				由于4与5互相指，执行head.Next=nil,4.Next=nil
				此时链表为1->2->3->4<-5
				返回节点5
		第三轮出栈：head=3,head.next=4
				执行head.next.next=head,4.next=3.此时链表为1->2->3<->4<-5
				由于3和4互指，执行head.next=nil,3.next=nil。此时链表为1->2->3<-4<-5
				返回节点5
		第四轮出栈：head=2,head.next=3
				执行head.next.next=head,3.next=2.此时链表为1->2<->3<-4<-5
				由于2和3互指，执行head.next=nil,2.next=nil。此时链表为1->2<-3<-4<-5
				返回节点5
		第五轮出栈：head=1,head.next=2
				执行head.next.next=head,2.next=1.此时链表为1<->2<-3<-4<-5
				由于1和2互指，执行head.next=nil,1.next=nil。此时链表为1<-2<-3<-4<-5
				返回节点5
		结束：最终返回5->4->3->2->1->nil
	*/
	head.Next.Next = head
	head.Next = nil

	return newHead
}
