import "strings"

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 *
 * https://leetcode.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (30.16%)
 * Total Accepted:    338.5K
 * Total Submissions: 1.1M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * Given a string, determine if it is a palindrome, considering only
 * alphanumeric characters and ignoring cases.
 *
 * Note: For the purpose of this problem, we define empty string as valid
 * palindrome.
 *
 * Example 1:
 *
 *
 * Input: "A man, a plan, a canal: Panama"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "race a car"
 * Output: false
 *
 *
 */
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l <= r {
		if !((s[l] >= 48 && s[l] <= 57) || (s[l] >= 'a' && s[l] <= 'z')) {
			l++
		} else if !((s[r] >= 48 && s[r] <= 57) || (s[r] >= 'a' && s[r] <= 'z')) {
			r--
		} else if s[l] == s[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

