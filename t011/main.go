package main

import (
	"fmt"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/133499
func main() {
	fmt.Printf("%v", solution([]string{
		"ayaye", "uuu", "yeye", "yemawoo", "ayaayaa",
	}))
}

func solution(babbling []string) int {
	var sum int
	for _, v := range babbling {
		v = strings.Replace(v, "aya", "1", -1)
		v = strings.Replace(v, "ye", "2", -1)
		v = strings.Replace(v, "woo", "3", -1)
		v = strings.Replace(v, "ma", "4", -1)
		var c int32
		for i, vv := range v {
			if c == vv ||
				(vv != 49 && vv != 50 && vv != 51 && vv != 52) {
				break
			}
			c = vv
			if i == len(v)-1 {
				sum++
			}
		}
	}
	return sum
}
