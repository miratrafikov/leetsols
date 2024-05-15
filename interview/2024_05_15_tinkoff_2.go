package main

import "fmt"

type Range struct {
	a, b int
}

func main() {
	var ranges []Range
	ranges = []Range{
		{a:0, b:1},
		{a:0, b:2},
		{a:1, b:3},
		{a:1, b:4},
	}
	_ = []Range{
		{a:1, b:2},
		{a:2, b:4},
		{a:3, b:6},
	}
	fmt.Println(countOverlappingRanges(ranges))
}

func countOverlappingRanges(ranges []Range) int {
	// сортируем диапазоны по неубыванию начал (левых краев)
	quicksort(ranges, 0, len(ranges)-1)
	// самый удаленный от нуля конец диапазона, при достижении которого прекращается прохождение по оси времени
	maxX := ranges[0].b
	// мапа точек к количеству диапазонов, оканчивающихся в этих точках
	m := make(map[int]int)
	// наполняем мапу, находим maxX
	for _, r := range ranges {
		m[r.b]++
		if r.b > maxX {
			maxX = r.b
		}
	}
	// текущий диапазон
	r := 0
	// текущее количество диапазонов, общих для текущей позиции на оси времени
	rangesAtX := 1
	// максимальное пересечение
	maxOverlap := 0
	// проходим по оси времени от ближайшего к нулю началу диапазона до maxX
	for x := ranges[0].a; x < maxX; x++ {
		// ищем количество диапазонов, начинающихся в X, и увеличиваем currOverlap на это количество, сдвигая r
		rangesStartingHere := 0
		for r + 1 < len(ranges) && ranges[r + 1].a == x {
			r++
			rangesStartingHere++
		}
		rangesAtX += rangesStartingHere
		// уменьшаем текущее пересечение на количество диапазонов, оканчивающихся в X
		if count, found := m[x]; found {
			rangesAtX -= count
			m[x] = 0
		}
		if rangesAtX > maxOverlap {
			maxOverlap = rangesAtX
		}
	}
	return maxOverlap
}

func quicksort(ranges []Range, l, r int) {
	if l == r {
		return
	}
	p := partition(ranges, l, r)
	quicksort(ranges, l, p)
	quicksort(ranges, p+1, r)
}

func partition(ranges []Range, l, r int) int {
	pivot := ranges[l+(r-l)/2]
	i := l - 1
	j := r + 1
	for {
		for {
			i++
			if !less(ranges[i], pivot) {
				break
			}
		}
		for {
			j--
			if !less(pivot, ranges[j]) {
				break
			}
		}
		if i >= j {
			break
		}
		ranges[i], ranges[j] = ranges[j], ranges[i]
	}
	return j
}

func less(a, b Range) bool {
	return a.a < b.a
}
