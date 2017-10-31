package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		_, ok := m[word]

		if ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}

	return m
}

func main() {
	fmt.Println(WordCount("Hello world Hello"))
}

/*
map[Hello:2 world:1]
*/
