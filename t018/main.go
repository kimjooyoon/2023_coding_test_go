package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%v", solution("12321", "42531"))
}

func split(s string) []string {
	var r = make([]string, len(s))
	for i, v := range s {
		r[i] = string(v)
	}
	return r
}

func solution(X string, Y string) string {
	var x = split(X)
	var r arr
	for i := range x {
		if strings.Contains(Y, x[i]) {
			Y = strings.Replace(Y, x[i], "", 1)
			r = append(r, x[i])
		}
	}
	var result string
	sort.Sort(r)
	for _, v := range r {
		result += v
	}
	if result == "" {
		return "-1"
	}
	tmp, _ := strconv.Atoi(result)
	if tmp == 0 {
		return "0"
	}

	return result
}

type arr []string

func (a arr) Len() int           { return len(a) }
func (a arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a arr) Less(i, j int) bool { return a[i] > a[j] }
