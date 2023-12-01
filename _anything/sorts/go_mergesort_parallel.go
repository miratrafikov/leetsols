package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

func main() {
	longArray := rand.Perm(10 * 100 * 100 * 100 * 10)
	longArrayCopy := make([]int, len(longArray))
	fmt.Println(longArray[:100])
	copy(longArrayCopy, longArray)
	var timeStart time.Time
	timeStart = time.Now()
	mergesortParallel(longArray, 0, len(longArray)-1)
	fmt.Println(time.Since(timeStart))
	timeStart = time.Now()	
	mergesort(longArrayCopy, 0, len(longArrayCopy)-1)
	fmt.Println(time.Since(timeStart))
}

func mergesort(arr []int, l, r int) {
	if l == r {
		return
	}
	m := l + (r - l) / 2
	mergesort(arr, l, m)
	mergesort(arr, m+1, r)
	merge(arr, l, m, r)
}

func mergesortParallel(arr []int, l, r int) {
	if l == r {
		return
	}
	m := l + (r - l) / 2
	var wg sync.WaitGroup
	wg.Add(2)
	go mergesortRoutine(arr, l, m, &wg)
	go mergesortRoutine(arr, m+1, r, &wg)
	merge(arr, l, m, r)
}

func mergesortRoutine(arr []int, l, r int, wg *sync.WaitGroup) {
	defer wg.Done()
	mergesortParallel(arr, l, r)
}

func merge(arr []int, l, m, r int) {
	sorted := make([]int, r-l+1)
	i, j, k := l, m+1, 0
	for i <= m && j <= r {
		if arr[i] < arr[j] {
			sorted[k] = arr[i]
			i++
		} else {
			sorted[k] = arr[j]
			j++
		}
		k++
	}
	for i <= m {
		sorted[k] = arr[i]
		i++
		k++
	}
	for j <= r {
		sorted[k] = arr[j]
		j++
		k++
	}
	i, j = l, 0
	for i <= r {
		arr[i] = sorted[j]
		i++
		j++
	}
}
