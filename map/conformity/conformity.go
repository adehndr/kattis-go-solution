package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	secondSolution()
}

func firstSolution() {
	scanner := bufio.NewScanner(os.Stdin)
	var limit int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &limit)
	aMap := make(map[int]int, limit)
	highest := 0
	for i := 0; i < limit; i++ {
		var a1, a2, a3, a4, a5 int
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d %d %d %d", &a1, &a2, &a3, &a4, &a5)
		sum := a1 + a2 + a3 + a4 + a5
		if v, ok := aMap[sum]; !ok {
			aMap[sum] = 1
		} else {
			v += 1
			aMap[sum] = v
		}
		if aMap[sum] >= highest {
			highest = aMap[sum]
		}
	}
	result := 0
	for _, v := range aMap {
		if v == highest {
			result += v
		}
	}
	fmt.Print(result)
}

func secondSolution() {
	scanner := bufio.NewScanner(os.Stdin)
	var limit int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &limit)
	aMap := map[string]int{}
	highest := 0
	for i := 0; i < limit; i++ {
		var a1, a2, a3, a4, a5 int
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d %d %d %d", &a1, &a2, &a3, &a4, &a5)
		tmpList := []int{a1, a2, a3, a4, a5}
		sort.Slice(tmpList, func(i, j int) bool { return tmpList[i] < tmpList[j] })
		// solved by convert the key to string
		var tmpString string
		for _, v := range tmpList {
			tmpString += fmt.Sprint(v)
		}

		if v, ok := aMap[tmpString]; !ok {
			aMap[tmpString] = 1
		} else {
			v += 1
			aMap[tmpString] = v
		}
		if aMap[tmpString] >= highest {
			highest = aMap[tmpString]
		}
	}
	result := 0
	for _, v := range aMap {
		if v == highest {
			result += v
		}
	}
	fmt.Print(result)
}
