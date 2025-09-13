package main

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)

	for _, name := range names {
		runes := []rune(name)
		if _, ok := counts[runes[0]]; !ok {
			counts[runes[0]] = make(map[string]int)
		}
		counts[runes[0]][name]++
	}

	return counts
}
