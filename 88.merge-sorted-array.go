/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 *
 * https://leetcode.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (34.80%)
 * Total Accepted:    340.6K
 * Total Submissions: 969.9K
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as
 * one sorted array.
 *
 * Note:
 *
 *
 * The number of elements initialized in nums1 and nums2 are m and n
 * respectively.
 * You may assume that nums1 has enough space (size that is greater or equal to
 * m + n) to hold additional elements from nums2.
 *
 *
 * Example:
 *
 *
 * Input:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 *
 * Output: [1,2,2,3,5,6]
 *
 *
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		nums1 = append(nums1[:m], nums2[:n]...)
		return
	}
	if n == 0 {
		return
	}
	for i := n - 1; i >= 0; i-- {
		last := nums1[m-1]
		j := m - 2
		for j >= 0 && nums1[j] > nums2[i] {
			nums1[j+1] = nums1[j]
			j--
		}
		if j != m-2 || last > nums2[i] {
			nums1[j+1] = nums2[i]
			nums2[i] = last
		}
	}
	nums1 = append(nums1[:m], nums2[:n]...)
}

