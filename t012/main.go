package main

import "fmt"

// https://school.programmers.co.kr/learn/courses/30/lessons/132267
// + 재귀로 풀어도 좋은듯
func main() {
	fmt.Printf("%v", solution(3, 1, 20))
}

func solution(a int, b int, n int) int {
	var result = 0

	for a <= n {
		tmp := b * (n / a)
		n = (n % a) + tmp
		result += tmp
	}

	return result
}
