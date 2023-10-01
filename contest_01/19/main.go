package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var input string
	var words []string

	for {
		fmt.Scan(&input)
		if input == "end" {
			break
		}
		words = append(words, input)
	}

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	var RepeatedWords []string
	for word, counter := range wordCount {
		if counter > 1 {
			RepeatedWords = append(RepeatedWords, word)
		}
	}

	sort.Strings(RepeatedWords)

	fmt.Println(strings.Join(RepeatedWords, " "))
}
