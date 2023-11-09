package main

import "fmt"

func main() {
	arr := []int{5,4,6,1,3,2}
	fmt.Println(arr)
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quicksort(arr []int, l, r int) {
	if l >= r {
		return
	}
	pivot := partition(arr, l, r)
	quicksort(arr, l, pivot)
	quicksort(arr, pivot+1, r)
}

func partition(arr []int, l, r int) int {
	pivot := arr[l + (r - l) / 2]
	i := l-1
	j := r+1
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
