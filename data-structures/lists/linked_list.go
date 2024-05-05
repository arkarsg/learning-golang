package lists

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
