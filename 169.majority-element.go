/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 *
 * https://leetcode.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (51.43%)
 * Total Accepted:    368.6K
 * Total Submissions: 706.8K
 * Testcase Example:  '[3,2,3]'
 *
 * Given an array of size n, find the majority element. The majority element is
 * the element that appears more than ⌊ n/2 ⌋ times.
 *
 * You may assume that the array is non-empty and the majority element always
 * exist in the array.
 *
 * Example 1:
 *
 *
 * Input: [3,2,3]
 * Output: 3
 *
 * Example 2:
 *
 *
 * Input: [2,2,1,1,1,2,2]
 * Output: 2
 *
 *
 */
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	m := len(nums) / 2
	e, count := 0, 0
	for _, v := range nums {
		t, ok := times[v]
		if !ok {
			times[v] = 1
			continue
		}
		if t+1 > m {
			return v
		}
		times[v] = t + 1
	}
	return 0
}

