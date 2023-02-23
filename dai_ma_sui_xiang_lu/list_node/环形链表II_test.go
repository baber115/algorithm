package list_node

import "testing"

// https://leetcode.cn/problems/linked-list-cycle-ii/description/
func TestDetectCycle(t *testing.T) {

}

// 快慢指针
/**
Q1：确认有环
快指针走两步，慢指针走一步，如果有环的话，必定在环中相遇
Q2：确认环的开始位置
设头节点到环入口的长度是x
环入口到相遇的点长度是y
相遇的点到环入口的长度是z
那么相遇时，慢指针走过的长度是：x + y，快指针走过的长度是x + y + n(y + z)，y + z是环的长度，n是在环里走了几圈
n一定是大于等于1的快指针一定是多走一圈才碰到慢指针
因为快指针一次走2步，慢指针一次走1步，所以他们走过的长度是：
2(x+y)=x+y+n(y+z)
两边消掉一个(x+y)：x+y=n(y+z)
因为求环的入口，也就是x的值，x=n(y+z)-y
从n里拿出一个(y+z)：x=(n-1)(y+z)+z
此时可以代入数字来计算，更方便
假设快指针走了1圈就碰上了慢指针，n=1，
此时x=z，可以得出头节点到环入口的长度 = 相遇的点到环入口的长度
这就意味着，从头结点出发一个指针，从相遇节点 也出发一个指针，这两个指针每次只走一个节点，
那么当这两个指针相遇的时候就是环形入口的节点。
*/
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	// 确认是环形链表，否则返回nil
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		// 计算出相遇的点
		if fast == slow {
			// 这就意味着，从头结点出发一个指针，从相遇节点 也出发一个指针，这两个指针每次只走一个节点，
			// 那么当这两个指针相遇的时候就是环形入口的节点。
			// 通过上面的论证过程，计算环的入口
			for slow != head {
				head = head.Next
				slow = slow.Next
			}
			return head
		}
	}

	return nil
}
