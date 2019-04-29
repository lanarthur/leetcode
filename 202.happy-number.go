/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */
func isHappy(n int) bool {
	if n == 0 {
		return false
	}
	numSum := func(x int) int {
		sum := 0
		for x != 0 {
			sum += (x % 10) * (x % 10)
			x /= 10
		}
		return sum
	}
	slow, fast := n, n
	for {
		slow = numSum(slow)
		fast = numSum(numSum(fast))
		if slow == fast {
			break
		}
	}
	return slow == 1
}

