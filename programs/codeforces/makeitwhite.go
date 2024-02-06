package main

import (
	"fmt"
	"math"
	"strings"
)

var word string
var length int
var T int
var firstOccr int
var secondOccr int

func returnAns(s []string) int {
	//for index, letter := range s {
	//	if letter == "B" && firstOccr != secondOccr {
	//		secondOccr = index
	//	}
	//	if letter == "B" {
	//		firstOccr = index
	//	}
	//	return int(math.Abs(float64(secondOccr)-float64(firstOccr))) + 1
	//}
	for i := 0; i < len(s); i++ {
		if s[i] == "B" {
			firstOccr = i
		}
	}
	for j := len(s) - 1; j >= 0; j-- {
		if s[j] == "B" {
			secondOccr = j
		}
	}
	return int(math.Abs(float64(secondOccr)-float64(firstOccr))) + 1
}
func main() {
	fmt.Scanln(&T)
	for i := 0; i < T; i++ {
		fmt.Scanln(&length)
		fmt.Scanln(&word)

		wordslice := strings.Split(word, "")
		ans := returnAns(wordslice)
		fmt.Println(ans)
		//	value, found := Map["B"]
		//	if found {
		//		ans = index - value
		//		break
		//	}
		//	Map[letter] = index
		//	ans = Map["B"]
		//	fmt.Println(ans)
		//
		//}
	}
}
