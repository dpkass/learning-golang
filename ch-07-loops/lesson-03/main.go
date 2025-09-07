package main

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 0
	balance := float64(maxCostInPennies)
	for balance >= actualCostInPennies {
		balance -= actualCostInPennies
		maxMessagesToSend++
		actualCostInPennies *= costMultiplier
	}
	return maxMessagesToSend
}
