package main

import (
	"fmt"
	"testing"
)

func TestInitSetGo(t *testing.T) {
	aSet := InitSet()

	aSet.Add(1)
	aSet.Add(1)
	aSet.Add(3)
	aSet.Add(4)
	aSet.Add(5)

	fmt.Println(aSet.Get())
}

const letters = "abcdefghijklmnopqrstuvwxyz"

func TestGenerate(t *testing.T) {
	lengthLetters := len(letters)
	for i := 1; i < 2; i++ {
		j := i
		var str string
		for j > 0 {
			str += string(letters[j%lengthLetters])
			j /= lengthLetters
		}
		fmt.Println(str)
	}
}

func TestMapWithinAMap(t *testing.T) {
	aMap := make(map[int]map[int]bool)
	aMap[0] = map[int]bool{0: true, 1: true}
	fmt.Println(aMap)
	if v, ok1 := aMap[1]; !ok1 {
		aMap[1] = map[int]bool{3: true}
	}else {
		if _, ok2 := v[3]; !ok2 {
			aMap[1][3] = true
		}
	}
	fmt.Println(aMap)

	bMap := map[int][]int{}
	if _, ok := bMap[0]; !ok {
		bMap[0] = []int{1}
	}else {
		bMap[0] = append(bMap[0], )
	}
}
