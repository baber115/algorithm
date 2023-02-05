package list_node

import (
	"testing"
)

// https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/
/**
指针相同
原来想法是翻转两个链表，然后比较，好像不行
*/
func TestIntersectionOfTwoLinkedLists(t *testing.T) {

}

// 哈希集合
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}

	for temp := headA; temp != nil; temp = temp.Next {
		vis[temp] = true
	}
	for temp := headB; temp != nil; temp = temp.Next {
		if vis[temp] {
			return temp
		}
	}

	return nil
}

// 双指针法一
/*
算出各自的长度，计算差值
长的比短的先走“差值”步
循环判断，如果一样就返回，不一样继续往后走
*/
func getIntersectionNodeI(headA, headB *ListNode) *ListNode {
	// 算出各自的长度，计算差值
	curA, curB := headA, headB
	headACount, headBCount := 0, 0
	for temp := headA; temp != nil; temp = temp.Next {
		headACount++
	}
	for temp := headB; temp != nil; temp = temp.Next {
		headBCount++
	}
	var step int
	var short, long *ListNode
	if headACount > headBCount {
		step = headACount - headBCount
		short, long = curB, curA
	} else {
		step = headBCount - headACount
		short, long = curA, curB
	}
	// 长的比短的先走“差值”步
	for i := 0; i < step; i++ {
		long = long.Next
	}

	// 判断指针是否相同
	for long != short {
		long = long.Next
		short = short.Next
	}

	return short
}

// 双指针法二，官方题解
/*
*
原理还是算差值
假设 listA = [4,1,8,4,5], listB = [5,0,1,8,4,5]
listBCount - listACount = 1
listA.Next循环结束，listA.next==nil，listB上的循环还差1步结束（此时差值产生）
然后listA指向listB，从listB重新开始循环，此时ListA指向ListB头节点
大家再走一步后
listB.Next循环结束，listB.next==nil，ListB指向ListA的头节点
由于上一步ListA已经指向ListB头节点，再走一步后，导致起始位置现在相同（也就是长的链表已经比短的链表多走了“差值”步）
*/
func getIntersectionNodeII(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	curA, curB := headA, headB

	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}

	return curA
}
