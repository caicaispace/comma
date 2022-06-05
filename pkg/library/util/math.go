package util

// 取两整数较小值
func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 取两整数较大值
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
