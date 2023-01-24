package t001

import "strings"

func solution(babbling []string) int {

	var i int
	for _, v := range babbling {
		v = strings.Replace(v, "aya", "-", 1)
		v = strings.Replace(v, "ye", "-", 1)
		v = strings.Replace(v, "woo", "-", 1)
		v = strings.Replace(v, "ma", "-", 1)
		v = strings.Replace(v, "-", "", -1)
		if v == "" {
			i = i + 1
		}
	}
	return i
}
