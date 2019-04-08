/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 *
 * https://leetcode.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (41.96%)
 * Total Accepted:    193.4K
 * Total Submissions: 452.6K
 * Testcase Example:  '3'
 *
 * Given a non-negative index k where k ≤ 33, return the k^th index row of the
 * Pascal's triangle.
 *
 * Note that the row index starts from 0.
 *
 *
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 *
 * Example:
 *
 *
 * Input: 3
 * Output: [1,3,3,1]
 *
 *
 * Follow up:
 *
 * Could you optimize your algorithm to use only O(k) extra space?
 *
 */
func getRow(rowIndex int) []int {
	value := 1
	row := rowIndex
	d := 1
	var tmp []int
	for j := 0; j <= rowIndex; j++ {
		tmp = append(tmp, value)
		value = value * row / d
		row--
		d++
	}
	return tmp
}

