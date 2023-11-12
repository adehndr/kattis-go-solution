package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	secondSolution()
}

func firstSolution() {
	scanner := bufio.NewScanner(os.Stdin)
	var limit int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &limit)
	var bitInput int
	for i := 0; i < limit; i++ {
		innerLimit := limit - i - 1
		scanner.Scan()
		switch scanner.Text() {
		case "Z":
			bitInput = bitInput | 0
		default:
			bitInput = bitInput | int(math.Pow(2, float64(innerLimit)))
		}
	}
	counterResult := 0
	for bitInput|0 != 0 {
		// loop from the least bit to most bit (right to left)
		// using the input limit as the limit
		// also, create the bit representation of 1 times how long the position
		// example
		// 1000 -> on position will be on the forth position onPosition = 3
		// but we create 111
		var onPosition int
		var multiplier int = 0
		for i := 0; i < limit; i++ {
			if bitInput&(1<<i) != 0 {
				onPosition = i
				break
			}
			multiplier = multiplier ^ (1 << i)
		}
		bitInput = bitInput ^ (1 << onPosition)
		bitInput = bitInput | multiplier
		counterResult += 1
	}
	fmt.Print(counterResult)
}

func printBinary(input int) {
	fmt.Println(strconv.FormatInt(int64(input), 2))
}

func secondSolution() {
	scanner := bufio.NewScanner(os.Stdin)
	var limit int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &limit)
	var output int
	for i := limit-1; i >= 0 ; i-- {
		scanner.Scan()
		if scanner.Text() == "O" {
			output = output ^ (1 << i)
		}
	}
  fmt.Print(output)
}