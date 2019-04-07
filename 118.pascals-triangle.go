/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 *
 * https://leetcode.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (44.47%)
 * Total Accepted:    238.8K
 * Total Submissions: 528.7K
 * Testcase Example:  '5'
 *
 * Given a non-negative integer numRows, generate the first numRows of Pascal's
 * triangle.
 *
 *
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 *
 * Example:
 *
 *
 * Input: 5
 * Output:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
 *
 *
 */
func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	var ret [][]int
	for i := 0; i < numRows; i++ {
		value := 1
		row := i
		d := 1
		var tmp []int
		for j := 0; j <= i; j++ {
			tmp = append(tmp, value)
			value = value * row / d
			row--
			d++
		}
		ret = append(ret, tmp)
	}
	return ret
}

