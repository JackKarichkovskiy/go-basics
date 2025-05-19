package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	var result = make(map[string]int)
	for _, word := range words {
		result[word]++
	}
	return result
}

func WordCountMain() {
	wordsCount := WordCount("I am a student I am a student I am student")
	fmt.Println(wordsCount)
}