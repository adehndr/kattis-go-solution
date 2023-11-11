package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// edge cases
// when input only calc. such as calc a =

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	mapWordToInt := map[string]int{}
	mapIntToWord := map[int]string{}
	for scanner.Scan() {
		line := scanner.Text()
		tmpList := strings.Split(line, " ")
		// when defining the key and value, don't forget to delete the old one
		// def ADE 12
		// def ADE 15
		// if using this logic, it will create a 2 key
		// ADE -> 12, 12 -> ADE
		// so if ade contains in the word to int,
		// delete the ADE -> 12, and 12 -> ADE
		if strings.Contains(line, "def") {
			resultConv, _ := strconv.ParseInt(tmpList[2], 10, 64)
			tempValueMapWordToInt, ok := mapWordToInt[tmpList[1]]
			if ok {
				delete(mapWordToInt, tmpList[1])
			}
			if _, ok := mapIntToWord[tempValueMapWordToInt]; ok {
				delete(mapIntToWord, tempValueMapWordToInt)
			}

			mapIntToWord[int(resultConv)] = tmpList[1]
			mapWordToInt[tmpList[1]] = int(resultConv)
		} else if strings.Contains(line, "calc") {
			var innerResult int
			var innerOp string
			var flag bool = true
		out:
			for i := 0; i < len(tmpList[1:])-1; i++ {
				if i%2 == 0 {
					// check if unknown or not
					var tempV int
					tempV, ok := mapWordToInt[tmpList[1+i]]
					if !ok {
						flag = false
						break out
					}
					if i == 0 {
						innerResult = int(tempV)
					} else {
						switch innerOp {
						case "+":
							innerResult += int(tempV)
						default:
							innerResult -= int(tempV)
						}
						innerOp = ""
					}
				} else {
					innerOp = tmpList[1+i]
				}
			}
			lastOutput := strings.Join(tmpList[1:], " ")
			if flag {
				// if result not exist
				if v, ok := mapIntToWord[innerResult]; ok {
					lastOutput += " " + v
				} else {
					lastOutput += " unknown"
				}
			} else {
				// unknown
				lastOutput += " unknown"
			}
			fmt.Println(lastOutput)
		} else {
			// clear
			mapWordToInt = map[string]int{}
			mapIntToWord = map[int]string{}
		}
	}
}
