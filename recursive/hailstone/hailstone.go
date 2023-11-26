package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	firstSolution()
}

func firstSolution() {
	scanner := bufio.NewScanner(os.Stdin)
	var input int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &input)
	result := simpleRecursive(input)
	fmt.Println(result)
}

var globalHMap map[int]int = make(map[int]int)

func simpleRecursive(n int) int {
	sum := n
	if n == 1 {
		globalHMap[sum] = 1
		return sum
	} else if n%2 == 0 {
		if v, ok := globalHMap[n/2]; ok {
			sum += v
		} else {
			sum += simpleRecursive(n / 2)
			globalHMap[n/2] = sum
		}
		return sum
	} else {
		if v, ok := globalHMap[3*n+1]; ok {
			sum += v
		} else {
			sum += simpleRecursive(3*n + 1)
			globalHMap[3*n+1] = sum
		}
		return sum
	}
}
