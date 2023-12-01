package main

import "fmt"

func main() {
	arr := []int{0,1,9,2,8,3,7,4,6,5}
	k := 4
	fmt.Printf("Find %dth smallest element in array %v\n", k+1, arr)
	lomutoSlice, hoareSlice := make([]int, 10), make([]int, 10)
	copy(lomutoSlice, arr)
	copy(hoareSlice, arr)
	lomutoAnswer := quickselectLomuto(lomutoSlice, 0, 9, k)
	hoareAnswer := quickselectHoare(hoareSlice, 0, 9, k)
	fmt.Printf("Answer by Lomuto is %d. Array: %v\n", lomutoAnswer, lomutoSlice)
	fmt.Printf("Answer by Hoare is %d. Array: %v\n", hoareAnswer, hoareSlice)
}

func quickselectLomuto(s []int, l, r, k int) int {
	if l == r {
		return s[l]
	}
	partitionPoint := partitionLomuto(s, l, r)
	if partitionPoint > k {
		return quickselectLomuto(s, l, partitionPoint-1, k)
	} else if partitionPoint < k {
		return quickselectLomuto(s, partitionPoint+1, r, k)
	}
	return s[partitionPoint]
}

func quickselectHoare(s []int, l, r, k int) int {
	if l == r {
		return s[l]
	}
	partitionPoint := partitionHoare(s, l, r)
	if partitionPoint >= k {
		return quickselectHoare(s, l, partitionPoint, k)
	}
	return quickselectHoare(s, partitionPoint+1, r, k)
}

func partitionLomuto(s []int, l, r int) int {
	pivot := s[r]
	i := l - 1
	for j := l; j < r; j++ {
		if s[j] < pivot {
			i++
			s[i], s[j] = s[j], s[i]
		}
	}
	i++
	s[i], s[r] = s[r], s[i]
	return i
}

func partitionHoare(s []int, l, r int) int {
	pivot := s[l + (r - l) / 2]
	i := l - 1
	j := r + 1
	for {
		for {
			i++
			if s[i] >= pivot {
				break
			}
		}
		for {
			j--
			if s[j] <= pivot {
				break
			}
		}
		if i >= j {
			break
		}
		s[i], s[j] = s[j], s[i]
	}
	return j
}
