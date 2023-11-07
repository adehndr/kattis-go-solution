package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lineChars string
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%s", &lineChars)
	scanner.Scan()
	tempWords := scanner.Text()
	lineWords := strings.Split(tempWords, " ")
	if len(lineWords) != len(lineChars) {
		fmt.Print("False")
		return
	}
	// char to word
	aMap := map[string]string{}
	// word to char
	bMap := map[string]string{}

	// used to build a long word
	var aBuilder strings.Builder
	// used to build a long char
	var bBuilder strings.Builder
	
	for i, chr := range lineChars {
		chrStr := string(chr)
		aMap[chrStr] = lineWords[i]
	}

	for i, word := range lineWords {
		bMap[word] = string(lineChars[i])
	}

	for _, chr := range lineChars {
		chrStr := string(chr)
		aBuilder.WriteString(aMap[chrStr])
	}

	for _, word := range lineWords {
		bBuilder.WriteString(bMap[word])
	}

	if aBuilder.String() == strings.Join(lineWords,"") && bBuilder.String() == lineChars {
		fmt.Print("True")
	}else {
		fmt.Print("False")
	}
	
}
