package main

func maxMessages(thresh int) int {
	for i, c := 0, 0; ; i, c = i+1, c+100+i {
		if c > thresh {
			return i - 1
		}
	}
}
