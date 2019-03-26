/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (32.98%)
 * Total Accepted:    415.8K
 * Total Submissions: 1.3M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 *
 * If there is no common prefix, return an empty string "".
 *
 * Example 1:
 *
 *
 * Input: ["flower","flow","flight"]
 * Output: "fl"
 *
 *
 * Example 2:
 *
 *
 * Input: ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 *
 * Note:
 *
 * All given inputs are in lowercase letters a-z.
 *
 */
func longestCommonPrefix(arr []string) string {
	if len(arr) == 0 {
		return ""
	}
	if len(arr) == 1 {
		return arr[0]
	}
	min := len(arr[0])
	for _, v := range arr[1:] {
		if len(v) < min {
			min = len(v)
		}
	}
	core := func(tmp []string, prefix string, start, end int) bool {
		for _, v := range tmp {
			for i := start; i <= end; i++ {
				if v[i] != prefix[i] {
					return false
				}
			}
		}
		return true
	}
	var (
		low, high = 0, min - 1
		prefix    string
	)
	for low <= high {
		mid := low + (high-low)/2
		if core(arr, arr[0], low, mid) {
			prefix = prefix + arr[0][low:mid+1]
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return prefix
}

