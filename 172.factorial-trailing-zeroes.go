import "math"

/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */
func trailingZeroes(n int) int {
	if n < 5 {
		return 0
	}
	ret := 0
	for i := 1; int(math.Pow(5, float64(i))) <= n; i++ {
		ret = ret + n/int(math.Pow(5, float64(i)))
	}
	return ret
}

