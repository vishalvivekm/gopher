package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sum1 := make([]rune, len(s))
	sum2 := make([]rune, len(t))

	for i, char := range s {
		sum1[i] = char
		sum2[i] = rune(t[i])
	}

	sort.Slice(sum1, func(i, j int) bool { return sum1[i] < sum1[j] })
	sort.Slice(sum2, func(i, j int) bool { return sum2[i] < sum2[j] })

	return strings.EqualFold(string(sum1), string(sum2))
}
