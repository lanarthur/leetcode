import "math/bits"

/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */
func reverseBits(num uint32) uint32 {
	return bits.Reverse32(num)
}

