type element struct {
    Value int
    Distance int
}

// minheap
type heap []element

func heapify(elements []element) *heap {
    h := make(heap, len(elements))
    for i := 0; i < len(elements); i++ {
        h[i] = elements[i]
        h.siftUp(i)
    }
    return &h
}

func (h *heap) siftUp(index int) {
    parentIndex := (index - 1) / 2
    for index > 0 && (*h).less((*h)[index], (*h)[parentIndex]) {
        (*h)[index], (*h)[parentIndex] = (*h)[parentIndex], (*h)[index]
        index = parentIndex
        parentIndex = (index - 1) / 2
    }
}

func (h *heap) less(a, b element) bool {
    return a.Distance < b.Distance || a.Distance == b.Distance && a.Value < b.Value
}

func (h *heap) DeleteTop() element {
    top := (*h)[0]
    (*h)[0] = (*h)[len((*h)) - 1]
    (*h) = (*h)[:len((*h)) - 1]
    (*h).siftDown(0)
    return top
}

func (h *heap) siftDown(index int) {
    leftChildIndex := 2*index + 1
    rightChildIndex := 2*index + 2
    minChildIndex := index
    if leftChildIndex < len((*h)) && (*h).less((*h)[leftChildIndex], (*h)[minChildIndex]) {
        minChildIndex = leftChildIndex
    }
    if rightChildIndex < len((*h)) && (*h).less((*h)[rightChildIndex], (*h)[minChildIndex]) {
        minChildIndex = rightChildIndex
    }
    if minChildIndex == index {
        return
    }
    (*h)[index], (*h)[minChildIndex] = (*h)[minChildIndex], (*h)[index]
    (*h).siftDown(minChildIndex)
}

func findClosestElements(arr []int, k int, x int) []int {
    n := len(arr)
    elements := make([]element, n)
    for i := 0; i < n; i++ {
        elements[i].Value = arr[i]
        elements[i].Distance = arr[i] - x
        if elements[i].Distance < 0 {
            elements[i].Distance = -elements[i].Distance
        }
    }
    heap := heapify(elements)
    closestValues := make([]int, k)
    for i := 0; i < k; i++ {
        minInHeap := heap.DeleteTop()
        closestValues[i] = minInHeap.Value
    }
    quicksort(closestValues, 0, k-1)
    return closestValues
}

func quicksort(arr []int, l, r int) {
    if l >= r {
        return
    }
    p := partition(arr, l, r)
    quicksort(arr, l, p)
    quicksort(arr, p+1, r)
}

func partition(arr []int, l ,r int) int {
    pivot := arr[l + (r - l) / 2]
    i := l - 1
    j := r + 1
    for {
        for {
            i++
            if arr[i] >= pivot {
                break
            }
        }
        for {
            j--
            if arr[j] <= pivot {
                break
            }
        }
        if i >= j {
            break
        }
        arr[i], arr[j] = arr[j], arr[i]
    }
    return j
}

