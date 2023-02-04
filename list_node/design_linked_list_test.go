package list_node

import (
	"fmt"
	"testing"
)

func TestDesignLinkedList(t *testing.T) {
	index := 1
	obj := Constructor()
	param_1 := obj.Get(index)
	obj.AddAtHead(1)
	obj.AddAtTail(2)
	obj.AddAtIndex(index, 3)
	obj.DeleteAtIndex(index)

	fmt.Println(param_1)
}

type MyLinkedList struct {
	dummy *Node
}

type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}

func Constructor() MyLinkedList {
	rear := &Node{
		Val:  -1,
		Next: nil,
		Pre:  nil,
	}
	rear.Next = rear
	rear.Pre = rear

	return MyLinkedList{rear}
}

func (this *MyLinkedList) Get(index int) int {
	return 0
}

func (this *MyLinkedList) AddAtHead(val int) {

}

func (this *MyLinkedList) AddAtTail(val int) {

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

}

func (this *MyLinkedList) DeleteAtIndex(index int) {

}
