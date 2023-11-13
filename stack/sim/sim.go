package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type stack struct {
	leftData  []string
	rightData []string
	pointer   bool // false -> left, true -> right
}

func (s *stack) delete() {
	if s.pointer {
		if len(s.rightData) == 0 {
			return
		}
		s.rightData = s.rightData[:len(s.rightData)-1]
	} else {
		if len(s.leftData) == 0 {
			return
		}
		s.leftData = s.leftData[:len(s.leftData)-1]
	}
}

func (s *stack) push(input string) {
	if s.pointer {
		s.rightData = append(s.rightData, input)
	} else {
		s.leftData = append(s.leftData, input)
	}
}

func (s *stack) rightP() {
	s.pointer = true
}

func (s *stack) leftP() {
	s.pointer = false
}

func initStack() *stack {
	return &stack{
		leftData:  make([]string, 0),
		rightData: make([]string, 0),
		pointer:   true,
	}
}

func (s *stack) merge() {
	s.rightData = append(s.leftData, s.rightData...)
	s.leftData = make([]string, 0)
}

func (s *stack) joining() string {
	total := s.leftData
	total = append(total, s.rightData...)
	return strings.Join(total, "")
}

const operator = "[]<"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var limit int
	fmt.Sscanf(scanner.Text(), "%d", &limit)
	for i := 0; i < limit; i++ {
		scanner.Scan()
		inputText := scanner.Text()
		aStack := initStack()
		for _, v := range inputText {
			if strings.Contains(operator, string(v)) {
				switch string(v) {
				case "[":
					aStack.merge()
					aStack.leftP()
				case "]":
					aStack.merge()
					aStack.rightP()
				default:
					aStack.delete()
				}
			} else {
				aStack.push(string(v))
			}
		}
		fmt.Println(aStack.joining())
	}
}
