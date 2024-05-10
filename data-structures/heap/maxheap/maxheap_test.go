package maxheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap(t *testing.T) {
	t.Run("Create a MaxHeap from array", func(t *testing.T) {
		arr := GetTestArray(t)
		heap := InitializeMaxHeap(t)
		heap.Create(arr)
		assert.True(
			t,
			HasHeapProperty(t, heap),
			"MaxHeap property violated: %v",
			heap.heap,
		)
	})

	t.Run("Inserting a larger value maintains MaxHeap", func(t *testing.T) {
		arr := GetTestArray(t)
		value := 100

		heap := InitializeMaxHeap(t)
		heap.Create(arr)
		heap.Insert(value)
		assert.True(
			t,
			HasHeapProperty(t, heap),
			"Inserting %v violated MaxHeap property: %v",
			value, heap.heap,
		)
	})

	t.Run("ExtractMax returns max value", func(t *testing.T) {
		arr := GetTestArray(t)
		value := 50

		heap := InitializeMaxHeap(t)
		heap.Create(arr)
		actual := heap.ExtractMax()
		assert.Equal(
			t,
			value,
			actual,
			"ExtractMax should extract %v, got %v",
			value, actual,
		)
	})

	t.Run("ExtractMax maintains MaxHeap property", func(t *testing.T) {
		arr := GetTestArray(t)
		initial := make([]int, len(arr))

		heap := InitializeMaxHeap(t)
		heap.Create(arr)
		copy(initial, heap.heap)
		_ = heap.ExtractMax()
		assert.True(
			t,
			HasHeapProperty(t, heap),
			"ExtractMax from %v violates MaxHeap property: %v",
			initial, heap.heap,
		)
	})
}

func HasHeapProperty(t testing.TB, h *MaxHeap) bool {
	t.Helper()
	n := len(h.heap)

	for i := 0; i <= (n/2)-1; i++ {
		leftChildIndex := i*2 + 1
		rightChildIndex := i*2 + 2

		if leftChildIndex < n && h.heap[leftChildIndex] > h.heap[i] {
			return false
		}

		if rightChildIndex < n && h.heap[rightChildIndex] > h.heap[i] {
			return false
		}
	}
	return true
}

func InitializeMaxHeap(t testing.TB) *MaxHeap {
	t.Helper()
	maxHeap := &MaxHeap{}
	return maxHeap
}

func GetTestArray(t testing.TB) []int {
	t.Helper()
	l := []int{2, 5, 10, 15, 19, 20, 30, 50}
	return l
}
