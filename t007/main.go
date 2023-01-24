package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("\n%v", solution(10, 3, 2))
}

func m(n int) int {
	var count int
	var i = 1
	for int(math.Sqrt(float64(n)))+1 > i {
		if n%i == 0 {
			if n/i == i {
				count = count + 1
			} else {
				count = count + 2
			}
		}
		i++
	}
	return count
}

func solution(number int, limit int, power int) int {
	var sum int
	for i := 1; i < number+1; i++ {
		if i == 1 {
			sum = sum + 1
		} else if m(i) > limit {
			sum = sum + power
		} else {
			sum = sum + m(i)
		}
	}
	return sum
}
