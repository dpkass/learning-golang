package main

func countConnections(groupSize int) (count int) {
	for i := 2; i <= groupSize; i++ {
		count += i - 1
	}
	return
}
