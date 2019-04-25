/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */
func rotate(nums []int, k int) {
	l := len(nums)
	if l == 0 || k == 0 {
		return
	}
	index := 0
	distance := 0
	cur := nums[0]
	for i := 0; i < l; i++ {
		index = (index + k) % l
		nums[index], cur = cur, nums[index]

		distance = (distance + k) % l
		if distance == 0 {
			index = (index + 1) % l
			cur = nums[index]
		}
	}
	return
}

