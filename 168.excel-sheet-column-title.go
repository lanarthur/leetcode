/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 *
 * https://leetcode.com/problems/excel-sheet-column-title/description/
 *
 * algorithms
 * Easy (28.49%)
 * Total Accepted:    168.4K
 * Total Submissions: 585.2K
 * Testcase Example:  '1'
 *
 * Given a positive integer, return its corresponding column title as appear in
 * an Excel sheet.
 *
 * For example:
 *
 *
 * ⁠   1 -> A
 * ⁠   2 -> B
 * ⁠   3 -> C
 * ⁠   ...
 * ⁠   26 -> Z
 * ⁠   27 -> AA
 * ⁠   28 -> AB
 * ⁠   ...
 *
 *
 * Example 1:
 *
 *
 * Input: 1
 * Output: "A"
 *
 *
 * Example 2:
 *
 *
 * Input: 28
 * Output: "AB"
 *
 *
 * Example 3:
 *
 *
 * Input: 701
 * Output: "ZY"
 *
 */
func convertToTitle(n int) string {
	var s []rune
	for n > 0 {
		rem := n % 26
		if rem == 0 {
			s = append(s, 'Z')
			n = n/26 - 1
		} else {
			s = append(s, rune(rem-1+int('A')))
			n = n / 26
		}
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

