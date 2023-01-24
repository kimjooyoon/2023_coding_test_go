package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v",
		solution(3, 4,
			[]int{1, 2, 3, 1, 2, 3, 1}),
	)
}

type scores []int

func (r scores) Len() int           { return len(r) }
func (r scores) Less(i, j int) bool { return r[i] > r[j] }
func (r scores) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func solution(k int, m int, score []int) int {
	var scoreArr scores = score
	sort.Sort(scoreArr)

	var result int
	for i := 0; i < len(score)/m; i++ {
		result += scoreArr[(i+1)*m-1] * m
	}

	return result
}
