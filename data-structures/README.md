# Data structures in Go

## LinkedList

> Implementation in `linked_list`

Singly linked list that holds an integer value:

```go
type Node struct {
	Value int
	Next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	length int
```

- `NewLinkedList()`
- `FromArray(arr []int)`
- `Append(value int)`
- `Prepend(value int)`
- `Search(value int)`
- `Get(index int)`
- `Insert(index int, value int)`
- `Remove(index int)`

### Tests

LinkedList tests are in `linked_list_test.go`

---

## MaxHeap

> Implementation in `heap`

MaxHeap data structure with integer keys

```go
/**
 * Zero indexed MaxHeap
 */
type MaxHeap struct {
	heap []int
}
```

- `Insert(value int)`
- `Create(arr []int)`
- `FastCreate(arr []int)`
- `ExtractMax()`
- `HeapSort()`
- `UpdateValue(target int, newValue int)`
- `DeleteValue(targetValue int)`

### Tests

Tests are in `maxheap_test.go`

---
