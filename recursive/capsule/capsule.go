package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	firstSolution()
}

var mapPosToRegion map[string]int = make(map[string]int)
var R int
var C int

func firstSolution() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d %d", &R, &C)
	matrix := make([][]int, 0)
	for i := 0; i < R; i++ {
		innerMatrix := make([]int, C)
		matrix = append(matrix, innerMatrix)
	}
	for i := 0; i < R; i++ {
		scanner.Scan()
		inputMatrix := scanner.Text()
		inputMatrixSplitted := strings.Split(inputMatrix, " ")
		for j, eachInput := range inputMatrixSplitted {
			if eachInput == "-" {
				matrix[i][j] = 0
			} else {
				res, _ := strconv.ParseInt(eachInput, 10, 64)
				matrix[i][j] = int(res)
			}
		}
	}
	for scanner.Scan() {
		input := scanner.Text()
		inputSpliited := strings.Split(input, " ")
		for _, pos := range inputSpliited[1:] {
			region, _ := strconv.ParseInt(inputSpliited[0], 10, 64)
			mapPosToRegion[pos] = int(region)
		}
	}
	fmt.Println(matrix)
	populateMatrix(matrix)
	fmt.Println(matrix)
}

// index-0 = row
// index-1 = col
var limitMatrix [][]int = [][]int{{0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}}

func isMatrixValid(matrix [][]int, row int, col int) bool {
	// validate surrounding number
	for i, limit := range limitMatrix {
		switch i {
		case 0:
			if row-1 >= 0 && matrix[row][col] == matrix[row+limit[0]][col+limit[1]] {
				fmt.Println("kiri aman tapi nilainya sama")
				return false
			}
		case 1:
			if row-1 >= 0 && col-1 >= 0 && matrix[row][col] == matrix[row+limit[0]][col+limit[1]] {
				fmt.Println("kiri atas aman tapi nilainya sama")
				return false
			}
		case 2:
			if col-1 >= 0 && matrix[row][col] == matrix[row+limit[0]][col+limit[1]] {
				fmt.Println("atas aman tapi nilainya sama")
				return false
			}
		case 3:
			if col+1 < C && row-1 >= 0 && matrix[row][col] == matrix[row+limit[0]][col+limit[1]] {
				fmt.Println("kanan atas aman tapi nilainya sama")
				return false
			}
		case 4:
			if col+1 < C && matrix[row][col] == matrix[row+limit[0]][col+limit[1]] {
				fmt.Println("kanan aman tapi nilainya sama")
				return false
			}
		case 5:
			if col+1 < C && row+1 < R && matrix[row][col] == matrix[row+limit[0]][col+limit[1]] {
				fmt.Println("kanan bawah aman tapi nilainya sama")
				return false
			}
		case 6:
			if row+1 < R && matrix[row][col] == matrix[row+limit[0]][col+limit[1]] {
				fmt.Println("bawah aman tapi nilainya sama")
				return false
			}
		default:
			if row+1 < R && col-1 >= 0 && matrix[row][col] == matrix[row+limit[0]][col+limit[1]] {
				fmt.Println("kiri bawah aman tapi nilainya sama")
				return false
			}
		}
	}

	// validate inner matrix

	return true
}

func populateMatrix(matrix [][]int) bool {
	for row := 0; row < R; row++ {
		for col := 0; col < C; col++ {
			if matrix[row][col] == 0 {
				for i := 1; i < 5; i++ {
					matrix[row][col] = i
					if isMatrixValid(matrix, row, col) && populateMatrix(matrix) {
						return true
					}
					matrix[row][col] = 0
				}
				return false
			}
		}
	}
	return true
}

func getValueFromMap(row int, col int) int {
	buildString := fmt.Sprintf("(%d,%d)", row, col)
	return mapPosToRegion[buildString]
}
