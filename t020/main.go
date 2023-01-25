package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%v", solution("12321", "42531"))
}

type StrHeap []string

func (h StrHeap) Len() int {
	return len(h)
}

func (h StrHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h StrHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StrHeap) Push(element interface{}) {
	*h = append(*h, element.(string))
}

func (h *StrHeap) Pop() interface{} {
	old := *h
	n := len(old)
	element := old[n-1]
	*h = old[0 : n-1]
	return element
}

func solution(X string, Y string) string {
	h := &StrHeap{}
	heap.Init(h)
	y := map[string]int{}
	for _, v := range Y {
		y[string(v)]++
	}

	for i := range X {
		if strings.Contains(Y, string(X[i])) {
			if y[string(X[i])] > 0 {
				y[string(X[i])]--
				heap.Push(h, string(X[i]))
			}
		}
	}

	var r = make([]string, h.Len())
	var i = 0
	for h.Len() > 0 {
		z := heap.Pop(h)
		r[i] = z.(string)
		i++
	}
	result := strings.Join(r, "")
	if result == "" {
		return "-1"
	}
	tmp, _ := strconv.Atoi(result)
	if tmp == 0 {
		return "0"
	}
	return result
}
