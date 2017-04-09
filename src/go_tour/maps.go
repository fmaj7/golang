package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCounts := make(map[string]int)
	for _, word := range words {
		if _, ok := wordCounts[word]; !ok {
			wordCounts[word] = 0
		}
		wordCounts[word] += 1
	}
	return wordCounts
}

func main() {
	wc.Test(WordCount)
}
