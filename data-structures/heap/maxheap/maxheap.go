package maxheap

/**
 * Zero indexed MaxHeap
 */
type MaxHeap struct {
	heap []int
}

func (h *MaxHeap) heapifyDown(index int) {
	largest := index
	leftChildIndex := 2*index + 1
	rightChildIndex := 2*index + 2
	n := len(h.heap)

	if leftChildIndex < n && h.heap[leftChildIndex] > h.heap[index] {
		largest = leftChildIndex
	}

	if rightChildIndex < n && h.heap[rightChildIndex] > h.heap[largest] {
		largest = rightChildIndex
	}

	if largest != index {
		h.heap[index], h.heap[largest] = h.heap[largest], h.heap[index]
		h.heapifyDown(largest)
	}
}

func (h *MaxHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.heap[index] <= h.heap[parentIndex] {
			break
		}
		h.heap[parentIndex], h.heap[index] = h.heap[index], h.heap[parentIndex]
		index = parentIndex
	}
}

func (h *MaxHeap) Insert(value int) {
	h.heap = append(h.heap, value)
	h.heapifyUp(len(h.heap) - 1)
}

func (h *MaxHeap) Create(arr []int) {
	for _, value := range arr {
		h.Insert(value)
	}
}

func (h *MaxHeap) ExtractMax() int {
	if len(h.heap) == 0 {
		panic("Heap is empty")
	}

	// save max value
	ret := h.heap[0]
	last := h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	// Remove the max value
	if len(h.heap) > 0 {
		h.heap[0] = last
		h.heapifyDown(0)
	}
	return ret
}
