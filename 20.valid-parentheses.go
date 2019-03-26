/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (35.86%)
 * Total Accepted:    525.7K
 * Total Submissions: 1.5M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 *
 *
 * Note that an empty string is also considered valid.
 *
 * Example 1:
 *
 *
 * Input: "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: "(]"
 * Output: false
 *
 *
 * Example 4:
 *
 *
 * Input: "([)]"
 * Output: false
 *
 *
 * Example 5:
 *
 *
 * Input: "{[]}"
 * Output: true
 *
 *
 */
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	var (
		arr []string
		x   string
	)
	for _, v := range s {
		strV := string(v)
		if (strV == "(") || (strV == "[") || (strV == "{") {
			arr = append(arr, strV)
			continue
		}
		if len(arr) == 0 {
			return false
		}
		switch strV {
		case ")":
			x, arr = arr[len(arr)-1], arr[:len(arr)-1]
			if (x == "{") || (x == "[") {
				return false
			}
			break
		case "]":
			x, arr = arr[len(arr)-1], arr[:len(arr)-1]
			if (x == "(") || (x == "{") {
				return false
			}
			break
		case "}":
			x, arr = arr[len(arr)-1], arr[:len(arr)-1]
			if (x == "(") || (x == "[") {
				return false
			}
			break
		}
	}
	if len(arr) != 0 {
		return false
	}
	return true
}

