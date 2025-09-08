package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, bw := range badWords {
			if word == bw {
				return i
			}
		}
	}
	return -1
}
