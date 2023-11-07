package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	mapWordToInt := map[string]int{}
	mapIntToWord := map[int]string{}
	for scanner.Scan() {
		line := scanner.Text()
		tmpList := strings.Split(line, " ")
		if strings.Contains(line, "def") {
			resultConv, _ := strconv.ParseInt(tmpList[2], 10, 64)
			mapIntToWord[int(resultConv)] = tmpList[1]
			mapWordToInt[tmpList[1]] = int(resultConv)
		} else if strings.Contains(line, "calc") {
			validateInput := []string{}
			for i, v := range tmpList {
				if i > 0 && i%2 != 0 {
					if _, ok := mapWordToInt[v]; ok {
						validateInput = append(validateInput,
							strconv.FormatInt(int64(mapWordToInt[v]), 10))
					} else {
						validateInput = append(validateInput, "-1001")
					}
				} else if i > 0 && i%2 == 0 {
					validateInput = append(validateInput, v)
				}
			}
			flag := false
			for _, v := range validateInput {
				if v == "-1001" {
					fmt.Println(strings.Join(tmpList[1:], " "), "unknown")
					flag = true
				}
			}
			if flag {
				continue
			}
			var result int
			for i, v := range validateInput {
				if i == 0 {
					res, _ := strconv.ParseInt(v, 10, 64)
					result = int(res)
				} else if i%2 == 0 {
					res, _ := strconv.ParseInt(v, 10, 64)
					switch validateInput[i-1] {
					case "+":
						result += int(res)
					default:
						result -= int(res)
					}
				}
			}
			if v, ok := mapIntToWord[result];ok {
				fmt.Println(strings.Join(tmpList[1:], " "), v)
			}else {
				fmt.Println(strings.Join(tmpList[1:], " "), "unknown")
			}
		} else {
			mapWordToInt = map[string]int{}
			mapIntToWord = map[int]string{}
		}
	}
}
