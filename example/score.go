package main

import (
	"fmt"
	"strings"
)

type LatterValue map[string]int

func Score(word string) int {
	latterMap := LatterValue{}
	latterMap["A"] = 1
	latterMap["A"] = 1
	latterMap["E"] = 1
	latterMap["I"] = 1
	latterMap["O"] = 1
	latterMap["U"] = 1
	latterMap["L"] = 1
	latterMap["N"] = 1
	latterMap["R"] = 1
	latterMap["S"] = 1
	latterMap["T"] = 1

	latterMap["D"] = 2
	latterMap["G"] = 2

	latterMap["B"] = 3
	latterMap["C"] = 3
	latterMap["M"] = 3
	latterMap["P"] = 3

	latterMap["F"] = 4
	latterMap["H"] = 4
	latterMap["V"] = 4
	latterMap["W"] = 4
	latterMap["Y"] = 4

	latterMap["K"] = 5

	latterMap["J"] = 8
	latterMap["X"] = 8

	latterMap["Q"] = 10
	latterMap["Z"] = 10

	sum := 0
	for i := 0; i < len(word); i++ {
		str := strings.ToUpper(string(word[i]))
		value := latterMap[str]
		sum += value
	}
	return sum
}

func main() {
	score := Score("abcdefghijklmnopqrstuvwxyz")
	fmt.Println("Score :", score)
}
