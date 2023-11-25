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
	for scanner.Scan() {
		input := scanner.Text()
		counter1 := 0
		counter0 := 0
		for _, aCh := range input {
			aChStr := fmt.Sprintf("%b", aCh)
			inner1 := 0
			for _, aChRune := range aChStr {
				if aChRune == 49 {
					counter1 += 1
					inner1 += 1
				}
			}
			counter0 += 7 - inner1
		}
		if counter1%2 == 0 && counter0%2 == 0 {
			fmt.Println("free")
		} else {
			fmt.Println("trapped")
		}
	}
}
