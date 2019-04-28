/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	v1 := nums[0]
	if n == 1 {
		return v1
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	v2 := max(nums[0], nums[1])
	if n == 2 {
		return v2
	}

	maxVal := v2

	for i := 2; i < n; i++ {
		maxVal = max(nums[i]+v1, v2)
		v1 = v2
		v2 = maxVal
	}
	return maxVal
}

