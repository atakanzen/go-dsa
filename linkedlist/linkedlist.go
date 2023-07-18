package linkedlist

import "fmt"

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

func NewDoublyLinkedList(values ...interface{}) *DoublyLinkedList {
	newList := &DoublyLinkedList{}
	for _, val := range values {
		newList.Add(val)
	}

	return newList
}

func (sl *DoublyLinkedList) Add(value interface{}) {
	newNode := &Node{Value: value}

	if sl.Head == nil {
		sl.Head = newNode
		sl.Tail = newNode
	} else {
		newNode.Prev = sl.Tail
		sl.Tail.Next = newNode
		sl.Tail = newNode
	}

}

func (sl *DoublyLinkedList) Prepend(value interface{}) {
	newNode := &Node{Value: value}

	if sl.Head == nil {
		sl.Head = newNode
		sl.Tail = newNode
	} else {
		newNode.Next = sl.Head
		sl.Head.Prev = newNode
		sl.Head = newNode
	}
}

func (sl *DoublyLinkedList) Get(value interface{}) *Node {
	if sl.Head == nil {
		return nil
	}

	if sl.Head.Value == value {
		return sl.Head
	}

	if sl.Tail != nil && sl.Tail.Value == value {
		return sl.Tail
	}

	n := sl.Head

	for n != nil {
		if value == n.Value {
			return n
		}
		n = n.Next
	}

	return nil
}

func (sl *DoublyLinkedList) Remove(value interface{}) bool {
	if sl.Head == nil {
		return false
	}

	if value == sl.Head.Value {
		sl.Head.Next.Prev = nil
		sl.Head = sl.Head.Next
	}

	if value == sl.Tail.Value {
		sl.Tail.Prev.Next = nil
		sl.Tail = sl.Tail.Prev
	}

	n := sl.Head

	for n != nil {
		if value == n.Value {
			n.Prev.Next, n.Next.Prev = n.Next, n.Prev
			return true
		}
		n = n.Next
	}

	return false
}

func (sl *DoublyLinkedList) Values() []interface{} {
	values := make([]interface{}, 0)

	if sl.Head != nil {
		n := sl.Head
		fmt.Println(sl.Tail)
		for n != nil {
			values = append(values, n.Value)
			n = n.Next
		}
	} else {
		return nil
	}

	return values
}
