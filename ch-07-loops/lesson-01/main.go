package main

func bulkSend(numMessages int) (cost float64) {
	cost = float64(numMessages)
	for i := 0; i < numMessages; i++ {
		cost += float64(i) / 100
	}
	return
}
