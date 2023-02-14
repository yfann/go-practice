package datastruct

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) InsertAtBeginning(data int) {
	node := &Node{data: data, next: ll.head}
	ll.head = node
}

func (ll *LinkedList) InsertAtEnd(data int) {
	node := &Node{data: data}
	if ll.head == nil {
		ll.head = node
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
}

func (ll *LinkedList) Delete(data int) {
	if ll.head == nil {
		return
	}
	if ll.head.data == data {
		ll.head = ll.head.next
		return
	}
	current := ll.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d ->", current.data)
		current = current.next
	}
	fmt.Println("nil")
}
