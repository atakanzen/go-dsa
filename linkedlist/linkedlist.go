package linkedlist

type SinglyLinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Value interface{}
	Next  *Node
}

func NewSinglyLinkedList(values ...interface{}) *SinglyLinkedList {
	newList := &SinglyLinkedList{}
	for _, val := range values {
		newList.Add(val)
	}

	return newList
}

func (sl *SinglyLinkedList) Add(value interface{}) {
	newNode := &Node{Value: value}

	if sl.Head == nil {
		sl.Head = newNode
		sl.Tail = newNode
	} else {
		sl.Tail.Next = newNode
		sl.Tail = newNode
	}

}

func (sl *SinglyLinkedList) Get(value interface{}) interface{} {
	n := sl.Head
	var found interface{}

	for n.Next != nil {
		if value == n.Value {
			found = n.Value
			break
		}
		n = n.Next
	}

	return found
}
