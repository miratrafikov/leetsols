// using quickselect
// since we need to find K most frequent elements in an array with N unique elements
// where N >= K, we need to run quickselect to find position of the N-K most frequent element
// and return most frequent elements starting N-K
type NumberAndFrequency struct {
    Number    int
    Frequency int
}
func topKFrequent(nums []int, k int) []int {
    numberToFrequency := make(map[int]int)
    for _, num := range nums {
        numberToFrequency[num] += 1
    }
    numbersWithFrequencies := make([]NumberAndFrequency, len(numberToFrequency))
    i := 0
    for number, frequency := range numberToFrequency {
        numbersWithFrequencies[i] = NumberAndFrequency{
            Number: number,
            Frequency: frequency,
        }
        i++
    }
    indexOfKthMostFrequentNumber := len(numbersWithFrequencies) - k
    quickselectNumberIndexByFrequency(
        numbersWithFrequencies,
        0,
        len(numbersWithFrequencies)-1,
        indexOfKthMostFrequentNumber,
    )
    kMostFrequentNumbers := make([]int, k)
    for i = 0; i < k; i++ {
        kMostFrequentNumbers[i] = numbersWithFrequencies[indexOfKthMostFrequentNumber+i].Number
    }
    return kMostFrequentNumbers
}

// will use Hoare's partition scheme therefore must keep partitioning until
// l is equal to r since even when m is equal to the target index it does not mean that
// elements to the right of the m are the largest in the array
func quickselectNumberIndexByFrequency(numbersWithFrequencies []NumberAndFrequency, l, r, k int) {
    if l == r {
        return
    }
    m := partition(numbersWithFrequencies, l, r)
    if m >= k {
        quickselectNumberIndexByFrequency(numbersWithFrequencies, l, m, k)
    } else if m < k {
        quickselectNumberIndexByFrequency(numbersWithFrequencies, m+1, r, k)
    }
}

func partition(numbersWithFrequencies []NumberAndFrequency, l, r int) int {
    pivot := numbersWithFrequencies[l + (r - l) / 2]
    i := l - 1
    j := r + 1
    for {
        for {
            i++
            if numbersWithFrequencies[i].Frequency >= pivot.Frequency {
                break
            }
        }
        for {
            j--
            if numbersWithFrequencies[j].Frequency <= pivot.Frequency {
                break
            }
        }
        if i >= j {
            break
        }
        numbersWithFrequencies[i], numbersWithFrequencies[j] = 
            numbersWithFrequencies[j], numbersWithFrequencies[i]
    }
    return j
}
