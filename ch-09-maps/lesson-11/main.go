package main

import "strings"

func countDistinctWords(messages []string) int {
	words := make(map[string]struct{})

	for _, message := range messages {
		if message == "" {
			continue
		}

		ws := strings.Split(message, " ")
		for _, w := range ws {
			w = strings.ToLower(w)
			words[w] = struct{}{}
		}
	}

	return len(words)
}
