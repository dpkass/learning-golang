package main

func getMessageWithRetries(primary, secondary, tertiary string) (messages [3]string, costs [3]int) {
	messages = [3]string{primary, secondary, tertiary}
	costs[0] = len(messages[0])
	costs[1] = costs[0] + len(messages[1])
	costs[2] = costs[1] + len(messages[2])
	return
}
