package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	for true {
		N := 0
		M := 0
		fmt.Fscan(in, &N, &M)
		if N == 0 && M == 0 {
			break
		}
		var result int
		aMap := make(map[int64]int64, 5e7)
		for i := 0; i < M+N; i++ {
			var input int64
			fmt.Fscan(in, &input)
			if _, ok := aMap[input]; !ok {
				aMap[input] = 1
			} else {
				result += 1
			}
		}
		fmt.Println(result)
	}
}
