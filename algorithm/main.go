package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var inputString string
	fmt.Scan(&inputString)

	pointMapper := getPointMapper(inputString)
	maxi := -1

	for first := 1; first < len(inputString); first++ {
		for second := first + 1; second < len(inputString); second++ {
			section1 := inputString[:first]
			section2 := inputString[first:second]
			section3 := inputString[second:]

			maxi = max(
				pointMapper[section1]+pointMapper[section2]+pointMapper[section3],
				maxi,
			)
		}
	}
	fmt.Println(maxi + 3)
}

func getPointMapper(inputString string) map[string]int {
	subStrings := make(map[string]struct{})
	getSubStringRecursive(0, inputString, subStrings)

	stringList := make([]string, 0, len(subStrings))
	for k := range subStrings {
		stringList = append(stringList, k)
	}
	sortStrings(stringList)

	pointMapper := make(map[string]int)
	for i, s := range stringList {
		pointMapper[s] = i
	}
	return pointMapper
}

func getSubStringRecursive(beginIndex int, inputString string, subStrings map[string]struct{}) {
	if beginIndex >= len(inputString) {
		return
	}
	for step := 1; step < len(inputString)-beginIndex; step++ {
		if beginIndex+step > len(inputString) {
			return
		}
		subStrings[inputString[beginIndex:beginIndex+step]] = struct{}{}
		getSubStringRecursive(beginIndex+step, inputString, subStrings)
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func sortStrings(strings []string) {
	for i := 0; i < len(strings)-1; i++ {
		for j := i + 1; j < len(strings); j++ {
			if strings[i] > strings[j] {
				strings[i], strings[j] = strings[j], strings[i]
			}
		}
	}
}
