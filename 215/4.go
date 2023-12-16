// max heap based solution
type MaxHeap struct {
    elements []int
    size     int
}

func newMaxHeap(expectedFinalSize int) MaxHeap {
    return MaxHeap{
        elements: make([]int, expectedFinalSize),
    }
}

func (h *MaxHeap) Add(v int) {
    index := h.size
    h.elements[index] = v
    h.siftUp()
    h.size++
}

func (h MaxHeap) siftUp() {
    index := h.size
    for index != 0 {
        parentIndex := (index - 1) / 2
        if h.elements[parentIndex] > h.elements[index] {
            break
        }
        h.swap(index, parentIndex)
        index = parentIndex
    }
}

func (h MaxHeap) swap(i, j int) {
    h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func (h *MaxHeap) PopMax() int {
    element := h.elements[0]
    h.elements[0] = h.elements[h.size - 1]
    h.siftDown()
    h.size--
    return element
}

func (h MaxHeap) siftDown() {
    index := 0
    for index < h.size {
        leftChildIndex := index * 2 + 1
        if leftChildIndex >= h.size {
            return
        }
        maxChildIndex := leftChildIndex
        rightChildIndex := index * 2 + 2
        if rightChildIndex < h.size && h.elements[rightChildIndex] > h.elements[maxChildIndex] {
            maxChildIndex = rightChildIndex
        }
        if h.elements[index] >= h.elements[maxChildIndex] {
            return
        }
        h.swap(index, maxChildIndex)
        index = maxChildIndex
    }
}

func findKthLargest(nums []int, k int) int {
    n := len(nums)
    maxHeap := newMaxHeap(n)
    for _, num := range nums {
        maxHeap.Add(num)
    }
    var kthLargestElement int
    for i := 0; i < k; i++ {
        kthLargestElement = maxHeap.PopMax()
    }
    return kthLargestElement
}
