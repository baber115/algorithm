package double_pointer

// https://leetcode.cn/problems/middle-of-the-linked-list/submissions/406048408/?envType=study-plan&id=suan-fa-ru-men&plan=algorithms&plan_progress=jgwt6s2&languageTags=golang
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil || fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
