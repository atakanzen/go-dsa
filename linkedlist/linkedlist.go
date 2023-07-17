package linkedlist

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
	n := sl.Head

	for n.Next != nil {
		if value == n.Value {
			return n
		}
		n = n.Next
	}

	return nil
}
