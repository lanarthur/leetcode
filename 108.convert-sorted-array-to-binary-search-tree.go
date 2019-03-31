/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
 *
 * https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/
 *
 * algorithms
 * Easy (49.17%)
 * Total Accepted:    246.7K
 * Total Submissions: 495.8K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * Given an array where elements are sorted in ascending order, convert it to a
 * height balanced BST.
 *
 * For this problem, a height-balanced binary tree is defined as a binary tree
 * in which the depth of the two subtrees of every node never differ by more
 * than 1.
 *
 * Example:
 *
 *
 * Given the sorted array: [-10,-3,0,5,9],
 *
 * One possible answer is: [0,-3,9,-10,null,5], which represents the following
 * height balanced BST:
 *
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
 *
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traverse(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	m := &TreeNode{Val: nums[mid]}
	m.Left = traverse(nums, start, mid-1)
	m.Right = traverse(nums, mid+1, end)
	return m
}
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return traverse(nums, 0, len(nums)-1)
}

