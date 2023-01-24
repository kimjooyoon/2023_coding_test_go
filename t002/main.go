package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//fmt.Printf("%v", solution("2022.05.19",
	//	[]string{"A 6", "B 12", "C 3"},
	//	[]string{"2021.05.02 A", "2021.07.01 B", "2022.02.19 C", "2022.02.20 C"},
	//))
	//
	//fmt.Printf("%v", solution("2020.01.01",
	//	[]string{"Z 3", "D 5"},
	//	[]string{"2019.01.01 D", "2019.11.15 Z", "2019.08.02 D", "2019.07.01 D", "2018.12.28 Z"},
	//))

	fmt.Printf("%v", solution("2019.01.01",
		[]string{"B 12"},
		[]string{"2017.12.21 B"},
	))
}

type date28 struct {
	y, m, d int
}

func toInt(s string) int {
	a, err := strconv.Atoi(s)
	if err != nil {
		panic("")
	}
	return a
}

func new28(y, m, d string, add int) date28 {

	d2, _ := strconv.Atoi(d)
	m2, _ := strconv.Atoi(m)
	m3 := m2 + add
	y2, _ := strconv.Atoi(y)
	if m3 > 12 {
		y2 = y2 + m3/12
		m3 = m3 - ((m3 / 12) * 12)
		if m3 < 1 {
			y2--
			m3 = m3 + 12
		}
	}

	return date28{
		y: y2,
		m: m3,
		d: d2,
	}
}

func (today date28) del(t date28) bool {
	if today.y < t.y {
		return false
	} else if today.y == t.y {
		if today.m < t.m {
			return false
		} else if today.m == t.m {
			if today.d < t.d {
				return false
			}
		}
	}

	return true
}

func split(today string) (y, m, d string) {
	return today[0:4], today[5:7], today[8:10]
}

type Terms []Term
type Term struct {
	key   string
	value int
}

func (arr Terms) FindByKey(key string) int {
	for _, term := range arr {
		if term.key == key {
			return term.value
		}
	}
	panic("")
	return 0
}

func solution(today string, terms []string, privacies []string) []int {
	y, m, d := split(today)
	todayD := new28(y, m, d, 0)

	var termArr = make(Terms, len(terms))

	for i, v := range terms {
		s := strings.Split(v, " ")
		termArr[i].key = s[0]
		termArr[i].value = toInt(s[1])
	}

	var dateArr = make([]date28, len(privacies))
	for i, privacy := range privacies {
		y, m, d := split(privacy)

		key := strings.Split(privacy, " ")[1]
		add := termArr.FindByKey(key)
		dateArr[i] = new28(y, m, d, add)
	}
	var result []int
	length := 0
	for i, v := range dateArr {
		if todayD.del(v) {
			length++
			result = append(result, i+1)
		}
	}

	var submit = make([]int, length)
	for i, v := range result {
		submit[i] = v
	}

	return submit
}
