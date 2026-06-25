func canJump(nums []int) bool {
    jumpable := make([]bool, len(nums))
	jumpable[0] = true

	for i, dist := range nums {
		if jumpable[i] {
			for jumpDist := range dist {
				if (i + jumpDist + 1) >= len(nums)-1 {
					return true
				}
				jumpable[i + jumpDist + 1] = true
			}
		}
	}

	return jumpable[len(nums)-1]
}
