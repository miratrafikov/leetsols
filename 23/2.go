// 96.93 time. using min heap to collecct all values, then dump into list
type minHeap []int

func newMinHeap() *minHeap {
    return &minHeap{}
}

func (h *minHeap) Push(val int) {
    *h = append(*h, val)
    h.siftUp()
}

func (h minHeap) siftUp() {
    index := len(h) - 1
    for index > 0 {
        parentIndex := (index - 1) / 2
        if h[parentIndex] <= h[index] {
            break
        }
        h.swap(index, parentIndex)
        index = parentIndex
    }
}

func (h *minHeap) PopMin() int {
    topElement := (*h)[0]
    lastElementIndex := len(*h) - 1
    (*h)[0] = (*h)[lastElementIndex]
    *h = (*h)[:lastElementIndex]
    h.siftDown()
    return topElement
}

func (h minHeap) siftDown() {
    index := 0
    heapSize := len(h)
    for index < heapSize - 1 {
        var minChildIndex int
        leftChildIndex := index * 2 + 1
        if leftChildIndex >= heapSize {
            return
        }
        minChildIndex = leftChildIndex
        rightChildIndex := index * 2 + 2
        if rightChildIndex < heapSize {
            if h[rightChildIndex] < h[leftChildIndex] {
                minChildIndex = rightChildIndex
            }
        }
        if h[minChildIndex] >= h[index] {
            return
        }
        h.swap(index, minChildIndex)
        index = minChildIndex
    }
}

func (h minHeap) swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h minHeap) IsEmpty() bool {
    return len(h) == 0
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    minHeap := newMinHeap()
    for i := 0; i < len(lists); i++ {
        for lists[i] != nil {
            minHeap.Push(lists[i].Val)
            lists[i] = lists[i].Next
        }
    }
    head := &ListNode{}
    tail := head
    for !minHeap.IsEmpty() {
        tail.Next = &ListNode{Val: minHeap.PopMin()}
        tail = tail.Next
    }
    return head.Next
}
