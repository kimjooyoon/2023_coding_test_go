package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%v", solution("3141592", "271"))
}

func solution(t string, p string) int {
	num, _ := strconv.Atoi(p)

	length := len(p)
	t_len := len(t) - length + 1

	count := 0
	for i := 0; i < t_len; i++ {
		ii, _ := strconv.Atoi(t[i : i+length])
		if ii <= num {
			count++
		}
	}

	return count
}
