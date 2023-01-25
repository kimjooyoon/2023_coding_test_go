package main

func solution(survey []string, choices []int) string {
	count := map[string]int{
		"R": 0,
		"T": 0,
		"C": 0,
		"F": 0,
		"J": 0,
		"M": 0,
		"A": 0,
		"N": 0,
	}
	for i, choice := range choices {
		if choice > 4 {
			count[string(survey[i][1])] += choice - 4
		} else if choice < 4 {
			count[string(survey[i][0])] += 4 - choice
		}
	}

	var result string
	if count["R"] >= count["T"] {
		result = "R"
	} else {
		result = "T"
	}
	if count["C"] >= count["F"] {
		result += "C"
	} else {
		result += "F"
	}
	if count["J"] >= count["M"] {
		result += "J"
	} else {
		result += "M"
	}
	if count["A"] >= count["N"] {
		result += "A"
	} else {
		result += "N"
	}

	return result
}
