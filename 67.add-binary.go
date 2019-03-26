import (
	"strconv"
)

/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 *
 * https://leetcode.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (37.88%)
 * Total Accepted:    282.4K
 * Total Submissions: 739.2K
 * Testcase Example:  '"11"\n"1"'
 *
 * Given two binary strings, return their sum (also a binary string).
 *
 * The input strings are both non-empty and contains only characters 1 or 0.
 *
 * Example 1:
 *
 *
 * Input: a = "11", b = "1"
 * Output: "100"
 *
 * Example 2:
 *
 *
 * Input: a = "1010", b = "1011"
 * Output: "10101"
 *
 */
func addBinary(a string, b string) string {
	var result string
	var s int

	i := len(a) - 1
	j := len(b) - 1
	for i >= 0 || j >= 0 || s == 1 {
		if i >= 0 {
			if string(a[i]) == "1" {
				s += 1
			}
		}
		if j >= 0 {
			if string(b[j]) == "1" {
				s += 1
			}
		}
		result = strconv.Itoa(s%2) + result
		s = s / 2
		i--
		j--
	}
	return result
}

