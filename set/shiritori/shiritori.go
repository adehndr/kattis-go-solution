package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	aMap := map[string]bool{}
	var input int
	var lastInput string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &input)
	for i := 0; i < input; i++ {
		var shiritori string
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%s", &shiritori)
		if i > 0 && (lastInput[len(lastInput)-1] != shiritori[0]) {
			if i%2 == 0 {
				fmt.Println("Player 1 lost")
			} else {
				fmt.Println("Player 2 lost")
			}
			return
		}

		if _, ok := aMap[shiritori]; ok {
			if i%2 == 0 {
				fmt.Println("Player 1 lost")
			} else {
				fmt.Println("Player 2 lost")
			}
			return
		}
		aMap[shiritori] = true
		lastInput = shiritori
	}
	fmt.Println("Fair Game")
}
