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
			str += string(letters[j % lengthLetters])
			j /= lengthLetters
		}
		fmt.Println(str)
	}
}

func TestAppendArray(t *testing.T) {
	aTemp := []byte{byte(97), byte(32), byte(98)}
	fmt.Println(string(aTemp))
	str := "ade"
	strByte := str[2]
	fmt.Println(string(strByte))
}