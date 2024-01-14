package main

import (
	"fmt"
	"unsafe"
)

type Node struct {
	Value int
	Both uintptr
}

type XorLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (this *XorLinkedList) Add(value int) {
	node := &Node{Value: value}
	nodePtr := uintptr(unsafe.Pointer(node))

	if this.Size == 0 {
		this.Head = node
		this.Tail = node
	} else if this.Size == 1 {
		this.Head.Both = nodePtr
		this.Tail = node
		node.Both = uintptr(unsafe.Pointer(this.Head))
	} else {
		node.Both = uintptr(unsafe.Pointer(this.Tail))
		this.Tail.Both = this.Tail.Both ^ nodePtr
		this.Tail = node
	}

	this.Size++
}

func (list *XorLinkedList) Get(index int) int {
	if index == 0 {
		return list.Head.Value
	} else if index == list.Size - 1 {
		return list.Tail.Value
	}

	prev := (*Node)(unsafe.Pointer(list.Head))
	current := (*Node)(unsafe.Pointer(list.Head.Both))

	for i := 1; i < index; i++ {
		tmp := current
		current = (*Node)(unsafe.Pointer(current.Both ^ uintptr(unsafe.Pointer(prev))))
		prev = tmp
	}

	return current.Value
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	list := XorLinkedList{}

	for _, v := range arr {
		list.Add(v)
	}

	for i := 0; i < len(arr); i++ {
		if list.Get(i) != arr[i] {
			fmt.Println("Failed")
		}	
	}

	fmt.Println("Success")
}