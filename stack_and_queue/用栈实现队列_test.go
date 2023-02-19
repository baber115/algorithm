package greedy_algorithm

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/implement-queue-using-stacks/

func TestImplementQueueUsingStacks(t *testing.T) {
	obj := QueueConstructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	fmt.Println(obj)
	param_2 := obj.Pop()
	param_3 := obj.Peek()
	param_4 := obj.Empty()
	fmt.Println(param_2, param_3, param_4)
}

type MyQueue struct {
	instack  []int
	outstack []int
}

func QueueConstructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.instack = append(this.instack, x)
}

func (this *MyQueue) handle() {
	for len(this.instack) > 0 {
		this.outstack = append(this.outstack, this.instack[len(this.instack)-1])
		this.instack = this.instack[:len(this.instack)-1]
	}
}

func (this *MyQueue) Pop() int {
	if len(this.outstack) == 0 {
		this.handle()
	}
	val := this.outstack[len(this.outstack)-1]
	this.outstack = this.outstack[:len(this.outstack)-1]

	return val
}

func (this *MyQueue) Peek() int {
	if len(this.outstack) == 0 {
		this.handle()
	}

	return this.outstack[len(this.outstack)-1]
}

func (this *MyQueue) Empty() bool {
	if len(this.outstack) == 0 {
		this.handle()
	}

	return len(this.outstack) == 0
}
