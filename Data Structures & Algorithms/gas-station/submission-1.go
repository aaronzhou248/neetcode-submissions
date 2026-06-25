func canCompleteCircuit(gas []int, cost []int) int {
    n := len(gas)

	startingLocation := 0
	currentGas := 0
	for i := 0; i < 2*n; i++ {
		if i - startingLocation >= n {
			return startingLocation
		}
		currentGas += gas[i % n]
		currentGas -= cost[i % n]

		if currentGas < 0 {
			currentGas = 0
			startingLocation = i + 1
		}
	}
	return -1
}
