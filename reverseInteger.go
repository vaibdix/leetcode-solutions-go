func reverse(x int) int {
    
	signMultiplier := 1
	if x < 0 {
		signMultiplier = -1
		x = signMultiplier * x
	}

	var result int
	for x > 0 {
		result = result*10 + x%10
		if signMultiplier*result >= math.MaxInt32 || signMultiplier*result <= math.MinInt32 {
			return 0
		}
		x = x / 10
	}
	return signMultiplier * result
}