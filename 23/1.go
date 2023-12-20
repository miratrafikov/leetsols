// using min heap
// slow as fuck, tle sometimes
// cot damn
type valAndIndex struct {
    Val   int
    Index int
}

type minHeap []valAndIndex

func newMinHeap() *minHeap {
    return &minHeap{}
}

func (h *minHeap) Push(val, index int) {
    *h = append(*h, valAndIndex{val, index})
    h.siftUp()
}

func (h minHeap) siftUp() {
    index := len(h) - 1
    for index > 0 {
        parentIndex := (index - 1) / 2
        if h[parentIndex].Val <= h[index].Val {
            break
        }
        h.swap(index, parentIndex)
        index = parentIndex
    }
}

func (h *minHeap) PopMin() (int, int) {
    topElement := (*h)[0]
    lastElementIndex := len(*h) - 1
    (*h)[0] = (*h)[lastElementIndex]
    *h = (*h)[:lastElementIndex]
    h.siftDown()
    return topElement.Val, topElement.Index
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
            if h[rightChildIndex].Val < h[leftChildIndex].Val {
                minChildIndex = rightChildIndex
            }
        }
        h.swap(index, minChildIndex)
        index = minChildIndex
    }
}

func (h minHeap) swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    tail := &ListNode{}
    head := tail
    k := len(lists)
    for k > 0 {
        minHeap := newMinHeap()
        didFindEmptyList := false
        for i := 0; i < k; i++ {
            if lists[i] == nil {
                lists = append(lists[:i], lists[i+1:]...)
                k = len(lists)
                didFindEmptyList = true
                break
            }
            minHeap.Push(lists[i].Val, i)
        }
        if didFindEmptyList {
            continue
        }
        smallestValue, listIndex := minHeap.PopMin()
        tail.Next = &ListNode{Val: smallestValue}
        tail = tail.Next
        lists[listIndex] = lists[listIndex].Next
    }
    return head.Next
}
