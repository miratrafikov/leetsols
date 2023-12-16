// using array-based priority queue. For educational purposes only
type NumberAndFrequency struct {
    Number    int
    Frequency int
}

type MaxHeap []NumberAndFrequency

func NewMaxHeap() MaxHeap {
    return make([]NumberAndFrequency, 0)
}

func (h *MaxHeap) Insert(e NumberAndFrequency) {
    *h = append(*h, e)
}

func (h *MaxHeap) DeleteMax() (NumberAndFrequency, error) {
    if len(*h) == 0 {
        return NumberAndFrequency{}, errors.New("heap is empty")
    }
    indexOfMax := 0
    for i := 0; i < len(*h); i++ {
        if (*h)[i].Frequency > (*h)[indexOfMax].Frequency {
            indexOfMax = i
        }
    }
    element := (*h)[indexOfMax]
    *h = append((*h)[:indexOfMax], (*h)[indexOfMax+1:]...)
    return element, nil
}

func topKFrequent(nums []int, k int) []int {
    numberToFrequency := make(map[int]int)
    for _, num := range nums {
        numberToFrequency[num] += 1
    }
    maxHeap := NewMaxHeap()
    for number, frequency := range numberToFrequency {
        maxHeap.Insert(NumberAndFrequency{
            Number: number,
            Frequency: frequency,
        })
    }
    kMostFrequentNumbers := make([]int, k)
    for i := 0; i < k; i++ {
        max, _ := maxHeap.DeleteMax()
        kMostFrequentNumbers[i] = max.Number
    }
    return kMostFrequentNumbers
}
