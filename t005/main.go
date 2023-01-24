package main

import "fmt"

func main() {
	fmt.Printf("%v", solution("abracadabra"))
}

func solution(s string) int {
	var count, c1, c2 int
	var word int32

	for i, v := range s {
		if word == 0 {
			word = v
		}
		if word == v {
			c1++
		} else {
			c2++
		}
		if c1 == c2 {
			count++
			word = 0
			continue
		}
		if i == len(s)-1 {
			count++
		}
	}
	return count
}
