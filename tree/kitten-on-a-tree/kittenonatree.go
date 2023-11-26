package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	value    int
	branches []node
}

// loop pada branches, dimana base case saat kondisi value equal dengan yg dicari
func (n *node) traverse() {
	for _, branch := range n.branches {
		fmt.Println("current :", branch.value)
		branch.traverse()
	}
}

var globalTrack []int

func (n *node) search(input int) {
	// get current parent node (root)
	// loop for every branch
	// do backtrack
	globalTrack = append(globalTrack, n.value)
	if n.branches == nil && n.value != input {
		globalTrack = globalTrack[:len(globalTrack)-1]
		return
	}

	for _, branch := range n.branches {
		branch.search(input)
	}

}

func main() {
	secondSolution()
}

func firstSolution() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var input int
	fmt.Sscanf(scanner.Text(), "%d", &input)
	hMapNode := make(map[int][]node)
	lengthTree := 0
	for scanner.Scan() {
		inputStr := scanner.Text()
		if inputStr == "-1" {
			break
		}
		lengthTree += 1
		arrInputStr := strings.Split(inputStr, " ")
		fmt.Println(arrInputStr)
		tempListNode := make([]node, len(arrInputStr)-1)
		for i, eachStr := range arrInputStr[1:] {
			tempValue, _ := strconv.ParseInt(eachStr, 10, 64)
			tempListNode[i] = node{
				value:    int(tempValue),
				branches: nil,
			}
		}
		tempRootNode, _ := strconv.ParseInt(arrInputStr[0], 10, 64)
		hMapNode[int(tempRootNode)] = tempListNode
	}
	for k, v := range hMapNode {
		for i, branch := range v {
			if tempV, ok := hMapNode[branch.value]; ok {
				hMapNode[k][i].branches = tempV
			}
		}
	}
}

func secondSolution() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var input int
	fmt.Sscanf(scanner.Text(), "%d", &input)
	hMap := make(map[int64]int64)
	for scanner.Scan() {
		inputStr := scanner.Text()
		if inputStr == "-1" {
			break
		}
		arrInputStr := strings.Split(inputStr, " ")
		tempRootNode, _ := strconv.ParseInt(arrInputStr[0], 10, 64)
		for _, eachNodeStr := range arrInputStr[1:] {
			tempBranchNode, _ := strconv.ParseInt(eachNodeStr, 10, 64)
			hMap[tempBranchNode] = tempRootNode
		}
	}
	trackTree(int64(input),hMap)
}

func trackTree(input int64, hMap map[int64]int64) {
	fmt.Print(input," ")
	if v, ok := hMap[input]; ok {
		trackTree(v,hMap)
	}else {
		return
	}
}
