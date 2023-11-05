package main

import (
	"bufio"
	"fmt"
	"os"
)

const lengthUnit = 2018

func main() {
	thirdSolution()
}

func firstSolution() {
	scanner := bufio.NewScanner(os.Stdin)
	var limit int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &limit)
	posS := make([][2]int, limit)
	count := 0
	for i := 0; i < limit; i++ {
		var x int
		var y int
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d", &x, &y)
		posS[i] = [2]int{x, y}
	}
	for i := 0; i < limit; i++ {
		for j := i + 1; j < limit; j++ {
			distance_x := posS[i][0] - posS[j][0]
			distance_y := posS[i][1] - posS[j][1]
			if distance_x*distance_x+distance_y*distance_y == lengthUnit*lengthUnit {
				count += 1
			}
		}
	}
	fmt.Print(count)
}

func correctDistance(xOri, yOri, xNext, yNext int) bool {
	correctPos := [][2]int{}
	for i := (-lengthUnit + xOri); i <= (lengthUnit + xOri); i++ {
		for j := (-lengthUnit + yOri); j <= (lengthUnit + yOri); j++ {
			if (j-yOri)*(j-yOri)+(i-xOri)*(i-xOri) == (lengthUnit)*(lengthUnit) {
				correctPos = append(correctPos, [2]int{i, j})
			}
		}
	}
	fmt.Println(correctPos)
	for _, v := range correctPos {
		if v[0] == xNext && v[1] == yNext {
			return true
		}
	}
	return false
}

func secondSolution() {
	scanner := bufio.NewScanner(os.Stdin)
	var limit int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &limit)
	posS := make([][2]int, limit)
	count := 0
	for i := 0; i < limit; i++ {
		var x int
		var y int
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d", &x, &y)
		posS[i] = [2]int{x, y}
	}

	for i := 0; i < len(posS)-1; i++ {
		if correctDistance(posS[i][0], posS[i][1], posS[i+1][0], posS[i+1][1]) {
			count += 1
		}
	}

	fmt.Print(count)
}

func generateCorrectList() [][2]int {
	result := make([][2]int, 12)
	counter := 0
	for i := -lengthUnit; i <= lengthUnit; i++ {
		for j := -lengthUnit; j <= lengthUnit; j++ {
			if i*i+j*j == lengthUnit*lengthUnit {
				result[counter] = [2]int{i, j}
				counter += 1
			}
		}
	}
	return result
}

func thirdSolution() {
	scanner := bufio.NewScanner(os.Stdin)
	var limit int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &limit)
	posS := make([][2]int, limit)
	count := 0
	for i := 0; i < limit; i++ {
		var x int
		var y int
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d", &x, &y)
		posS[i] = [2]int{x, y}
	}
	deltas := generateCorrectList()
	seen := make(map[int][]int, limit)
	for i := 0; i < limit; i++ {
		for _, delta := range deltas {
			// another pos will check the previous point
			// example 1680,1118 will check 0,0
			// C(3360,0) also will checked, but since C is not added to seen map
			// the counter only incremented by 1
			anotherPos := [2]int{posS[i][0] + delta[0], posS[i][1] + delta[1]}
			if listPoint, ok := seen[anotherPos[0]]; ok {
				for _, point := range listPoint {
					if point == anotherPos[1] {
						count += 1
					}
				}
			}
		}
		// previous point will be added in seen so on the next iteration
		// the previous point will be checked
		if _, ok := seen[posS[i][0]]; !ok {
			seen[posS[i][0]] = []int{posS[i][1]}
		} else {
			seen[posS[i][0]] = append(seen[posS[i][0]], posS[i][1])
		}
	}

	fmt.Print(count)
}
