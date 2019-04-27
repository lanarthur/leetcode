import "math/bits"

/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */
func hammingWeight(num uint32) int {
	return bits.OnesCount32(num)
}

