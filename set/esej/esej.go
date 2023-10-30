package main

import (
	"bufio"
	"fmt"
	"os"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var min int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d %d", &min, &min)
	var result string
	for i := 1; i <= min; i++ {
		result += generateWord(i) + " "
	}
	fmt.Print(result)
}

func generateWord(n int) string {
	lengthLetters := len(letters)
	var tmp string
	for n > 0 {
		tmp += string(letters[n%lengthLetters])
		n /= lengthLetters
	}
	return tmp
}

/* func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandString() string {
	lengthWord := rand.Intn(1) + 1
	b := make([]byte, lengthWord)
	for i := range b {
		b[i] = letters[rand.Intn(26)]
	}
	return string(b)
}
*/

/*
func failedMethod() {
	scanner := bufio.NewScanner(os.Stdin)
	var min int
	var max int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d %d", &min, &max)
	result := make(map[string]bool, min)
	var output string
	for len(result) < min {
		outputStr := RandString()
		if _, ok := result[outputStr]; !ok {
			result[outputStr] = true
			if len(result) < min {
				output += outputStr + " "
			} else {
				output += outputStr
			}
		}
	}
	fmt.Print(output)
}
*/
