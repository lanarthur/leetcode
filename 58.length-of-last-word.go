import "strings"

/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 *
 * https://leetcode.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (32.15%)
 * Total Accepted:    250.6K
 * Total Submissions: 778.9K
 * Testcase Example:  '"Hello World"'
 *
 * Given a string s consists of upper/lower-case alphabets and empty space
 * characters ' ', return the length of last word in the string.
 *
 * If the last word does not exist, return 0.
 *
 * Note: A word is defined as a character sequence consists of non-space
 * characters only.
 *
 * Example:
 *
 * Input: "Hello World"
 * Output: 5
 *
 *
 */
func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	arr := strings.Split(strings.TrimSpace(s), " ")
	return len(arr[len(arr)-1])
}

