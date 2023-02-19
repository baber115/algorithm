package greedy_algorithm

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/implement-stack-using-queues/

func TestImplementStackUsingQueues(t *testing.T) {
	obj := StackConstructor()
	obj.Push(1)
	obj.Push(2)
	param_3 := obj.Top()
	param_2 := obj.Pop()
	param_4 := obj.Empty()
	fmt.Println(param_2, param_3, param_4)
}

type MyStack struct {
	queue []int
}

func StackConstructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	size := len(this.queue) - 1
	for size > 0 {
		val := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, val)
		size--
	}
	val := this.queue[0]
	this.queue = this.queue[1:]
	return val
}

func (this *MyStack) Top() int {
	return this.queue[len(this.queue)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
