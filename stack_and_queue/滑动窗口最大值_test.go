package greedy_algorithm

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/sliding-window-maximum/

func TestMaxSlidingWindow(t *testing.T) {
	// 3, 3, 5, 5, 6, 7
	//nums, k := []int{1, 3, -1, -3, 5, 3, 6, 7}, 3
	nums, k := []int{1, -1}, 1
	res := maxSlidingWindow(nums, k)
	fmt.Println(res)
}

func maxSlidingWindow(nums []int, k int) []int {
	m := newMyQueue()
	var res []int
	for i := 0; i < k; i++ {
		m.push(nums[i])
	}
	res = append(res, m.front())

	for j := k; j < len(nums); j++ {
		m.pop(nums[j-k])
		m.push(nums[j])
		res = append(res, m.front())
	}

	return res
}

type myQueue struct {
	queue []int
}

func newMyQueue() *myQueue {
	return &myQueue{
		queue: make([]int, 0),
	}
}

func (m *myQueue) pop(val int) {
	if len(m.queue) > 0 && m.queue[0] == val {
		m.queue = m.queue[1:]

	}
}

func (m *myQueue) front() int {
	return m.queue[0]
}

func (m *myQueue) push(val int) {
	for len(m.queue) > 0 && val > m.queue[len(m.queue)-1] {
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)
}
