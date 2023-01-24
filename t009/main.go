package main

import (
	"fmt"
	"strconv"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/134240
func main() {
	fmt.Printf("%v", solution([]int{1, 3, 4, 6}))
}

func solution(food []int) string {
	var result string
	for i, v := range food {
		if i == 0 {
			continue
		}
		for j := 0; j < v/2; j++ {
			result = result + strconv.Itoa(i)
		}
	}
	tmp := reverse(result)
	result = result + "0"
	result = result + tmp

	return result
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}
