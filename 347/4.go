// using array-based heap. For educational purposes only
type NumberAndFrequency struct {
    Number    int
    Frequency int
}

type MaxHeap []NumberAndFrequency

func MaxHeapify(numberToFrequency map[int]int) MaxHeap {
    heap := make(MaxHeap, len(numberToFrequency))
    numberOfElementsAddedSoFar := 0
    for number, frequency := range numberToFrequency {
        numberAndFrequency := NumberAndFrequency{
            Number: number,
            Frequency: frequency,
        }
        heap[numberOfElementsAddedSoFar] = numberAndFrequency
        numberOfElementsAddedSoFar++
        heap.siftUp(numberOfElementsAddedSoFar-1)
    }
    return heap
}

func (h *MaxHeap) siftUp(i int) {
    parentIndex := (i - 1) / 2
    for i > 0 && (*h)[parentIndex].Frequency < (*h)[i].Frequency {
        (*h)[parentIndex], (*h)[i] = (*h)[i], (*h)[parentIndex]
        i = parentIndex
        parentIndex = (i - 1) / 2
    }
}

func (h *MaxHeap) DeleteMax() (NumberAndFrequency, error) {
    if len(*h) == 0 {
        return NumberAndFrequency{}, errors.New("heap is empty")
    }
    element := (*h)[0]
    (*h)[0] = (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    (*h).siftDown(0)
    return element, nil
}

func (h *MaxHeap) siftDown(i int) {
    leftChildIndex := 2*i + 1
    rightChildIndex := 2*i + 2
    largestChildIndex := i
    if leftChildIndex < len(*h) && (*h)[leftChildIndex].Frequency > (*h)[largestChildIndex].Frequency {
        largestChildIndex = leftChildIndex
    }
    if rightChildIndex < len(*h) && (*h)[rightChildIndex].Frequency > (*h)[largestChildIndex].Frequency {
        largestChildIndex = rightChildIndex
    }
    if largestChildIndex != i {
        (*h)[i], (*h)[largestChildIndex] = (*h)[largestChildIndex], (*h)[i]
        h.siftDown(largestChildIndex)
    }
}

func topKFrequent(nums []int, k int) []int {
    numberToFrequency := make(map[int]int)
    for _, num := range nums {
        numberToFrequency[num] += 1
    }
    maxHeap := MaxHeapify(numberToFrequency)
    kMostFrequentNumbers := make([]int, k)
    for i := 0; i < k; i++ {
        max, _ := maxHeap.DeleteMax()
        kMostFrequentNumbers[i] = max.Number
    }
    return kMostFrequentNumbers
}
