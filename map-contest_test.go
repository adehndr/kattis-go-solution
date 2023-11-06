package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortingArray(t *testing.T) {
	anArr := []int{2, 1, 3}
	sort.Slice(anArr, func(i, j int) bool {return anArr[i] < anArr[j]})
	fmt.Println(anArr)
}
