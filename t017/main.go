package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v", solution([]int{-1, -1, -1}))
	fmt.Printf("%v", o2([]int{-1, -1, -1}))
}

func o2(number []int) int {
	var result int
	l := len(number)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if i >= l || j >= l || k >= l {
					continue
				}
				if number[i]+number[j]+number[k] == 0 {
					result++
				}
			}
		}
	}

	return result
}

// https://school.programmers.co.kr/learn/courses/30/lessons/131705
func solution(number []int) int {
	var result int
	var nums arr = number
	var p1, p2, p3 = 0, 1, len(number) - 1

	sort.Sort(nums)

	for {
		if p1 >= nums.Len() {
			break
		}
		if nums[p1] <= 0 {
			break
		}
		p3 = len(number) - 1
		for {
			if nums[p3] >= 0 {
				break
			}
			p2 = p1 + 1
			for {
				if nums[p1]+nums[p2]+nums[p3] < 0 || p2 >= p3 {
					break
				} else if nums[p1]+nums[p2]+nums[p3] == 0 {

					//fmt.Printf("a:%v b:%v c:%v\n", nums[p1], nums[p2], nums[p3])
					//fmt.Printf("x:%v y:%v z:%v\n", p1, p2, p3)
					result++
				}
				p2++
			}
			p3--
		}
		p1++
	}
	zero := 0
	for _, v := range nums {
		if v == 0 {
			zero++
		}
	}

	if zero == 3 {
		result++
	} else if zero > 3 {
		tmp := 1
		for i := 4; i <= zero; i++ {
			tmp *= i
		}
		result += tmp
	}

	return result
}

type arr []int

func (a arr) Len() int           { return len(a) }
func (a arr) Less(i, j int) bool { return a[i] > a[j] }
func (a arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
