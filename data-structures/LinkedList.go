package data_structures

import "fmt"

type LinkedListNode struct {
	next *LinkedListNode
	data int
}

type LinkedList struct {
	Head *LinkedListNode
}

func NewLinkedList() *LinkedList  {
	return &LinkedList{Head: nil}
}

func (l *LinkedList)AddNode(data int) *LinkedListNode {
	node := &LinkedListNode{data: data}
	if l.Head == nil {
		l.Head = node
	} else {
		curr := l.Head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = node
	}

	return node
}

func (l *LinkedList) Print()  {
	if l.Head == nil {
		return
	} else {
		curr := l.Head
		for curr.next != nil {
			fmt.Println(curr.data)
			curr = curr.next
		}
		fmt.Println(curr.data)
	}
}
