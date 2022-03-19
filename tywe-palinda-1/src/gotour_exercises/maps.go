package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	word := strings.Fields(s)
	wordMap := make(map[string]int)
	for _, comp := range word {
		wordMap[comp]++
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
