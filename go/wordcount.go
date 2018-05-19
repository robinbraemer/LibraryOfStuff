package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int, len(s))
	for _, w := range strings.Fields(s) {
		if i, ex := m[w]; ex {
			m[w] = i + 1
		} else {
			m[w] = 1
		}
	}
	
	return m
}

func main() {
	fmt.Println()
	wc.Test(WordCount)
}
