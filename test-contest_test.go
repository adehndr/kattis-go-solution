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
