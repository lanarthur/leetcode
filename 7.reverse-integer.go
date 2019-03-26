import (
	"math"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (25.11%)
 * Total Accepted:    615.8K
 * Total Submissions: 2.5M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 *
 * Example 1:
 *
 *
 * Input: 123
 * Output: 321
 *
 *
 * Example 2:
 *
 *
 * Input: -123
 * Output: -321
 *
 *
 * Example 3:
 *
 *
 * Input: 120
 * Output: 21
 *
 *
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose
 * of this problem, assume that your function returns 0 when the reversed
 * integer overflows.
 *
 */
func reverse(x int) int {
	if x < math.MinInt32 || x > math.MaxInt32 {
		return 0
	}
	str := strings.Split(strconv.FormatInt(int64(x), 10), "")
	l := len(str)
	reverseStr := make([]string, len(str))
	for k, v := range str {
		if (v == "0" && k == l-1) || v == "-" {
			continue
		}
		reverseStr[l-k-1] = v
	}
	ret, err := strconv.ParseInt(strings.Join(reverseStr, ""), 10, 32)
	if err != nil {
		return 0
	}
	if x < 0 {
		ret = -ret
	}
	return int(ret)
}

