package main

import data_structures "github.com/theodesp/go-hackerearth/data-structures"

func main() {
	l := data_structures.NewLinkedList()
	l.AddNode(2)
	l.AddNode(18)
	l.AddNode(24)
	l.AddNode(3)
	l.AddNode(5)
	l.AddNode(7)
	l.AddNode(9)
	l.AddNode(6)
	l.AddNode(1)
	l.ReverseEvens()
	l.Print()
}
