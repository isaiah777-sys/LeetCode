package leetcode

type MedianFinder struct {
	minHeap []int
	maxHeap []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: []int{},
		maxHeap: []int{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.maxHeap) == 0 || num <= this.maxHeap[0] {
		this.maxHeap = append(this.maxHeap, num)
		this.maxHeap = maxHeapifyUp(this.maxHeap, len(this.maxHeap)-1)
	} else {
		this.minHeap = append(this.minHeap, num)
		this.minHeap = minHeapifyUp(this.minHeap, len(this.minHeap)-1)
	}
	if len(this.maxHeap) > len(this.minHeap)+1 {
		this.minHeap = append(this.minHeap, this.maxHeap[0])
		this.minHeap = minHeapifyUp(this.minHeap, len(this.minHeap)-1)
		this.maxHeap[0] = this.maxHeap[len(this.maxHeap)-1]
		this.maxHeap = this.maxHeap[:len(this.maxHeap)-1]
		this.maxHeap = maxHeapifyDown(this.maxHeap, 0)
	} else if len(this.minHeap) > len(this.maxHeap) {
		this.maxHeap = append(this.maxHeap, this.minHeap[0])
		this.maxHeap = maxHeapifyUp(this.maxHeap, len(this.maxHeap)-1)
		this.minHeap[0] = this.minHeap[len(this.minHeap)-1]
		this.minHeap = this.minHeap[:len(this.minHeap)-1]
		this.minHeap = minHeapifyDown(this.minHeap, 0)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.maxHeap) > len(this.minHeap) {
		return float64(this.maxHeap[0])
	}
	return float64(this.maxHeap[0]+this.minHeap[0]) / 2.0
}

func maxHeapifyUp(heap []int, index int) []int {
	for index > 0 {
		parent := (index - 1) / 2
		if heap[index] > heap[parent] {
			heap[index], heap[parent] = heap[parent], heap[index]
			index = parent
		} else {
			break
		}
	}
	return heap
}

func minHeapifyUp(heap []int, index int) []int {
	for index > 0 {
		parent := (index - 1) / 2
		if heap[index] < heap[parent] {
			heap[index], heap[parent] = heap[parent], heap[index]
			index = parent
		} else {
			break
		}
	}
	return heap
}

func maxHeapifyDown(heap []int, index int) []int {
	for index < len(heap) {
		left := 2*index + 1
		right := 2*index + 2
		largest := index
		if left < len(heap) && heap[left] > heap[largest] {
			largest = left
		}
		if right < len(heap) && heap[right] > heap[largest] {
			largest = right
		}
		if largest == index {
			break
		}
		heap[index], heap[largest] = heap[largest], heap[index]
		index = largest
	}
	return heap
}

func minHeapifyDown(heap []int, index int) []int {
	for index < len(heap) {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index
		if left < len(heap) && heap[left] < heap[smallest] {
			smallest = left
		}
		if right < len(heap) && heap[right] < heap[smallest] {
			smallest = right
		}
		if smallest == index {
			break
		}
		heap[index], heap[smallest] = heap[smallest], heap[index]
		index = smallest
	}
	return heap
}
