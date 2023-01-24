package main

import "fmt"

func main() {
	fmt.Printf("%v", solution(
		//[]int{1, 2, 1, 1, 1, 1, 1, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 3, 1}))
		[]int{1, 2, 1, 2, 3, 1, 3, 1, 1, 2, 3, 1}))
}

// https://school.programmers.co.kr/learn/courses/30/lessons/133502
// 보류
func solution(ingredient []int) int {
	var count int
	var c, c2 int
	var p = []int{1, 2, 3, 1}

	var attach = 1

	for i := 0; i < len(ingredient); i++ {
		if ingredient[i] == p[c] {
			if c == 3 {
				count++
				back := i - (4 * attach)
				if back >= 0 {
					i2 := back
					c2 = ingredient[i2] - 1
					tmp := c2

					for {
						if tmp == 0 {
							c = c2 + 1
							attach++
							break
						}

						if p[tmp] == ingredient[i2] {
							if i2 == 0 {
								attach = 1
								break
							}
							i2--
							tmp--
						} else {
							c = 0
							attach = 1
							break
						}
					}
				} else {
					c = 0
					attach = 1
				}
			} else {
				c++
			}
		} else if ingredient[i] == 1 {
			c = 1
			attach = 1
		} else {
			c = 0
			attach = 1
		}
	}
	return count
}
