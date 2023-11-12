package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type stack struct {
	data    []string
	pointer bool // false -> left, true -> right
}

func (s *stack) delete() {
	if len(s.data) == 0 {
		return
	}

	if s.pointer {
		s.data = s.data[:len(s.data)-1]
	} else {
		s.data = s.data[1:]
	}
}

func (s *stack) push(input string) {
	if s.pointer {
		s.data = append(s.data, input)
	} else {
		// should change this logic
		// original = "AB", append C and D
		// it should be CDAB not DCAB
		s.data = append([]string{input}, s.data...)
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
		data:    nil,
		pointer: true,
	}
}

func (s *stack) joining() string {
	return strings.Join(s.data, "")
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
					aStack.leftP()
				case "]":
					aStack.rightP()
				default:
					aStack.delete()
				}
			} else {
				aStack.push(string(v))
			}
		}
		fmt.Print(aStack.joining())
	}
}
