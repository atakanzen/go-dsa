package linkedlist

type DoublyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
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

func (dl *DoublyLinkedList) Add(value interface{}) {
	newNode := &Node{Value: value}

	if dl.Head == nil {
		dl.Head = newNode
		dl.Tail = newNode
	} else {
		newNode.Prev = dl.Tail
		dl.Tail.Next = newNode
		dl.Tail = newNode
	}

	dl.Length += 1
}

func (dl *DoublyLinkedList) Prepend(value interface{}) {
	newNode := &Node{Value: value}

	if dl.Head == nil {
		dl.Head = newNode
		dl.Tail = newNode
	} else {
		newNode.Next = dl.Head
		dl.Head.Prev = newNode
		dl.Head = newNode
	}

	dl.Length += 1
}

func (dl *DoublyLinkedList) Get(value interface{}) *Node {
	if dl.Head == nil {
		return nil
	}

	if dl.Head.Value == value {
		return dl.Head
	}

	if dl.Tail != nil && dl.Tail.Value == value {
		return dl.Tail
	}

	n := dl.Head

	for n != nil {
		if value == n.Value {
			return n
		}
		n = n.Next
	}

	return nil
}

func (dl *DoublyLinkedList) Remove(value interface{}) bool {
	if dl.Head == nil {
		return false
	}

	if value == dl.Head.Value {
		dl.Head.Next.Prev = nil
		dl.Head = dl.Head.Next
		dl.Length -= 1
		return true
	}

	if value == dl.Tail.Value {
		dl.Tail.Prev.Next = nil
		dl.Tail = dl.Tail.Prev
		dl.Length -= 1
		return true
	}

	n := dl.Head

	for n != nil {
		if value == n.Value {
			n.Prev.Next, n.Next.Prev = n.Next, n.Prev
			dl.Length -= 1
			return true
		}
		n = n.Next
	}

	return false
}

func (dl *DoublyLinkedList) RemoveAt(index int) bool {
	if dl.Length < 1 {
		return false
	}

	if index == 0 {
		return dl.Remove(dl.Head.Value)
	}

	if index == dl.Length-1 {
		return dl.Remove(dl.Tail.Value)
	}

	if index < dl.Length/2 { // Start from Head
		node := dl.Head
		for i := 0; i <= index; i++ {
			if i == index {
				// Remove
				node.Prev.Next, node.Next.Prev = node.Next, node.Prev
				dl.Length -= 1
				return true
			}
			node = node.Next
		}
	} else { // Start from Tail
		node := dl.Tail
		for i := dl.Length - 1; i >= index; i-- {
			if i == index {
				node.Prev.Next, node.Next.Prev = node.Next, node.Prev
				dl.Length -= 1
				return true
			}
			node = node.Prev
		}
	}

	return false
}

func (dl *DoublyLinkedList) Values() []interface{} {
	values := make([]interface{}, 0)

	if dl.Head != nil {
		n := dl.Head
		for n != nil {
			values = append(values, n.Value)
			n = n.Next
		}
	} else {
		return nil
	}

	return values
}
