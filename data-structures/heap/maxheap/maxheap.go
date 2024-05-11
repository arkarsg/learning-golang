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

func (h *MaxHeap) FastCreate(arr []int) {
	n := len(arr)
	h.heap = []int{}

	for i := 0; i < (n/2)-1; i++ {
		h.heap = append(h.heap, arr[i])
	}

	for i := (n / 2); i < n; i++ {
		h.heap = append(h.heap, arr[i])
		h.heapifyUp(len(h.heap) - 1)
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

func (h *MaxHeap) HeapSort() (sorted []int) {
	sorted = []int{}
	n := len(h.heap)
	for i := 0; i < n; i++ {
		val := h.ExtractMax()
		sorted = append(sorted, val)
	}
	return
}

func (h *MaxHeap) UpdateValue(targetValue int, newValue int) {
	for index, value := range h.heap {
		if value == targetValue {
			h.heap[index] = newValue
			if newValue > targetValue {
				h.heapifyUp(index)
			} else {
				h.heapifyDown(index)
			}
		}
	}
}

func (h *MaxHeap) DeleteValue(targetValue int) {
	for index := 0; index < len(h.heap); index++ {
		value := h.heap[index]
		if value == targetValue {
			h.heap[index] = h.heap[len(h.heap)-1]
			h.heap = h.heap[:len(h.heap)-1]
			h.heapifyDown(index)
		}
	}
}
