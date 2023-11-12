package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type boolStack struct {
	items []*bool
}

func newStack() *boolStack {
	return &boolStack{
		items: nil,
	}
}

func (bs *boolStack) Pop() *bool {
	if len(bs.items) == 0 {
		return nil
	}

	lastItem := bs.items[len((*bs).items)-1]
	bs.items = bs.items[:len(bs.items)-1]

	return lastItem
}

func (bs *boolStack) Push(input bool) {
	bs.items = append(bs.items, &input)
}

func (bs *boolStack) print() {
	for i, v := range bs.items {
		if i == len(bs.items)-1 {
			fmt.Println(*v)
		} else {
			fmt.Print(*v, " ")
		}
	}
}

const operator = "+-*"

func main() {
	// fmt.Println(strings.Contains(operator, "*"))
	firstSolution()
}

func firstSolution() {
	scanner := bufio.NewScanner(os.Stdin)
	var limit int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &limit)
	scanner.Scan()
	lineBoolStr := strings.Split(scanner.Text(), " ")
	lineBoolConverted := convertSliceStringToSliceBool(lineBoolStr)
	scanner.Scan()
	var inputGate string = scanner.Text()
	counterSliceBool := 0
	mapStringToBool := make(map[string]bool, len(lineBoolConverted))
	for _, v := range inputGate {
		if counterSliceBool == len(lineBoolConverted) {
			break
		}
		// if not +=* or space
		if !strings.Contains(operator, string(v)) && v != 32 {
			mappingStringToBool(&mapStringToBool, string(v), lineBoolConverted[counterSliceBool])
			counterSliceBool += 1
		}
	}

	aStack := newStack()
	for _, v := range inputGate {
		if v == 32 {
			continue
		}
		if strings.Contains(operator, string(v)) {
			if string(v) == "-" {
				aStack.Push(!(*(aStack.Pop())))
			} else {
				a1 := aStack.Pop()
				a2 := aStack.Pop()
				switch string(v) {
				case "+":
					aStack.Push(*a1 || *a2)
				default:
					aStack.Push(*a1 && *a2)
				}
			}
		} else if !strings.Contains(operator, string(v)) && v != 32 {
			aStack.Push(mapStringToBool[string(v)])
		}
	}
	finalRes := aStack.Pop()
	if *finalRes {
		fmt.Print("T")
	}else {
		fmt.Print("F")
	}
}

func convertSliceStringToSliceBool(input []string) []bool {
	result := make([]bool, len(input))
	for i, v := range input {
		res, _ := strconv.ParseBool(v)
		result[i] = res
	}
	return result
}

func mappingStringToBool(inputMap *map[string]bool, inputStr string, inputBool bool) {
	(*inputMap)[inputStr] = inputBool
}
