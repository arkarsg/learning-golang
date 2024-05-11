package maxheap

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	l := []int{2, 5, 10, 15, 15, 19, 20, 30, 50}
	return l
}

func GetDescendingSorted(t testing.TB) []int {
	t.Helper()
	l := GetTestArray(t)
	sort.Sort(sort.Reverse(sort.IntSlice(l)))
	return l
}

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

	t.Run("FastCreate creates a MaxHeap", func(t *testing.T) {
		arr := GetTestArray(t)
		heap := InitializeMaxHeap(t)
		heap.FastCreate(arr)
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

	t.Run("HeapSort creates sorted non-decreasing array", func(t *testing.T) {
		arr := GetTestArray(t)
		expected := GetDescendingSorted(t)

		heap := InitializeMaxHeap(t)
		heap.Create(arr)
		sortedArray := heap.HeapSort()
		assert.Equal(t, expected, sortedArray)
	})

	t.Run("Update value to a larger value maintains MaxHeap property", func(t *testing.T) {
		arr := GetTestArray(t)
		targetValue := 15
		newValue := 40

		heap := InitializeMaxHeap(t)
		heap.Create(arr)
		heap.UpdateValue(targetValue, newValue)
		assert.True(
			t,
			HasHeapProperty(t, heap),
		)
	})

	t.Run("Update value to a smaller value maintains MaxHeap property", func(t *testing.T) {
		arr := GetTestArray(t)
		targetValue := 15
		newValue := 3

		heap := InitializeMaxHeap(t)
		heap.Create(arr)
		heap.UpdateValue(targetValue, newValue)
		assert.True(
			t,
			HasHeapProperty(t, heap),
		)
	})

	t.Run("Delete value maintans MaxHeap property", func(t *testing.T) {
		arr := GetTestArray(t)
		targetValue := 15

		heap := InitializeMaxHeap(t)
		heap.Create(arr)
		heap.DeleteValue(targetValue)
		assert.True(
			t,
			HasHeapProperty(t, heap),
		)
	})
}
