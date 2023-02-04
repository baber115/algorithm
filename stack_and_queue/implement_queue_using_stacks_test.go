package stack_and_queue

import (
	"fmt"
	"testing"
)

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
func TestImplementQueueUsingStacks(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	peek := queue.Peek() // 返回 1
	fmt.Printf("peek: %d\n", peek)
	pop := queue.Pop() // 返回 1
	fmt.Printf("pop: %d\n", pop)
	empty := queue.Empty() // 返回 false
	fmt.Printf("empty: %t\n", empty)
}

type MyQueue struct {
	stack []int
	back  []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack: make([]int, 0),
		back:  make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	for len(this.back) != 0 {
		val := this.back[len(this.back)-1]
		this.back = this.back[:len(this.back)-1]
		this.stack = append(this.stack, val)
	}
	this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
	for len(this.stack) != 0 {
		val := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.back = append(this.back, val)
	}
	if len(this.back) == 0 {
		return 0
	}
	val := this.back[len(this.back)-1]
	this.back = this.back[:len(this.back)-1]
	return val
}

func (this *MyQueue) Peek() int {
	val := this.Pop()
	if val == 0 {
		return 0
	}
	this.back = append(this.back, val)
	return val
}

func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0 && len(this.back) == 0
}
