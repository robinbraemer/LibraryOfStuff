package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCountMap := make(map[string]int, len(s))
	
	for _, w := range words {
		wordCountMap[w]++
	}
	
	return wordCountMap
}

func main() {
	wc.Test(WordCount)
}
