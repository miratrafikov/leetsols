package main

import "fmt"

func main() {
	nums := []int{5,4,7,12,365,21,1,4,2,66}
	quicksort(&nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func quicksort(nums *[]int, l, h int) {
	if l < h && l >= 0 && h >= 0 {
		fmt.Println(*nums)
		p := partition(nums, l, h)
		quicksort(nums, l, p)
		quicksort(nums, p+1, h)
	}
}

func partition(nums *[]int, l, h int) int {
	p := (*nums)[l + (h-l)/2]
	i := l-1
	j := h+1
	for {
		for {
			i++
			if (*nums)[i] >= p {
				break
			}
		}
		for {
			j--
			if (*nums)[j] <= p {
				break
			}
		}
		if i >= j {
			break
		}
		(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
	}
	return j
}
