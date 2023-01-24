package main

import "fmt"

func main() {
	fmt.Printf("%v", solution("banana"))
}

type data struct {
	result int
	char   int32
}
type dataArr []data

func reverse(s []int32) []int32 {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
func solution(s string) []int {
	m := make(map[string]int, len(s))
	var res []int

	for i := range s {
		val, ok := m[string(s[i])]
		if !ok {
			m[string(s[i])] = i
			res = append(res, -1)
		} else {
			m[string(s[i])] = i
			res = append(res, i-val)
		}
	}
	return res
}
