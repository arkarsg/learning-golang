package linked_list

type Node struct {
	Value int
	Next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	length int
}

func Hello() string {
	return "Hello from lists"
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func FromArray(nodes []int) *LinkedList {
	l := NewLinkedList()
	for _, value := range nodes {
		l.Append(value)
	}
	return l
}

func (l *LinkedList) Length() int {
	return l.length
}

func (l *LinkedList) Head() *Node {
	return l.head
}

func (l *LinkedList) Tail() *Node {
	return l.tail
}

func (l *LinkedList) Append(value int) {
	addedNode := &Node{Value: value}
	if l.head == nil {
		l.head = addedNode
	} else {
		l.tail.Next = addedNode
	}
	l.tail = addedNode
	l.length++
}

func (l *LinkedList) Prepend(value int) {
	addedNode := &Node{Value: value}
	if l.head == nil {
		l.tail = addedNode
	} else {
		addedNode.Next = l.head
	}
	l.head = addedNode
	l.length++
}

func (l *LinkedList) Search(value int) int {
	if l.head == nil {
		return -1
	}

	idx := 0
	curr := l.head
	for curr != nil {
		if curr.Value == value {
			return idx
		}
		curr = curr.Next
		idx++
	}
	return -1
}

func (l *LinkedList) Get(idx int) *Node {
	if idx == 0 {
		return l.head
	}

	if idx == l.length - 1 {
		return l.tail
	}

	if idx >= l.length || idx < 0 {
		return nil
	}

	currNode := l.head
	for i := 0; i < idx; i++ {
		currNode = currNode.Next
	}
	return currNode
}

func (l *LinkedList) Insert(idx int, val int) {
	if idx < 0 || idx > l.Length() {
		return
	}

	if l.head == nil || idx == 0 {
		l.Prepend(val)
		return
	}

	if idx == l.Length() {
		l.Append(val)
		return
	}

	nodeAdded := &Node{Value: val}
	prev := l.Get(idx - 1)
	after := prev.Next
	nodeAdded.Next = after
	prev.Next = nodeAdded
	l.length++
}

func (l *LinkedList) Remove(idx int) {
	if idx < 0 || idx > l.Length() {
		return
	}

	if idx == 0 {
		tmp := l.head
		l.head = l.head.Next
		tmp.Next = nil
		l.length--
		return
	}

	prev := l.Get(idx - 1)
    if prev == nil || prev.Next == nil {
        return // Index out of bounds or previous node is nil, do nothing
    }

    if idx == l.length - 1 {
    	prev.Next = nil
     	l.tail = prev
        l.length--
        return
    }
}
