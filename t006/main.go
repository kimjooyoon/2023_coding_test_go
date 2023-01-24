package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v", solution(4,
		[]int{0, 300, 40, 300, 20, 70, 150, 50, 500, 1000}))
}

type sorter []int

func (s sorter) Len() int {
	return len(s)
}

func (s sorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sorter) Less(i, j int) bool {
	return s[i] > s[j]
}

func solution(k int, score []int) []int {
	var k_arr = sorter{}
	var result = make([]int, len(score))

	for i, v := range score {
		k_arr = append(k_arr, v)
		sort.Sort(k_arr)

		if i <= k-1 {
			result[i] = k_arr[i]
		} else {
			result[i] = k_arr[k-1]
			k_arr = k_arr[:k]
		}
	}

	return result
}
