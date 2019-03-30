/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (45.59%)
 * Total Accepted:    214.7K
 * Total Submissions: 466.9K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the bottom-up level order traversal of its
 * nodes' values. (ie, from left to right, level by level from leaf to root).
 *
 *
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 *
 * return its bottom-up level order traversal as:
 *
 * [
 * ⁠ [15,7],
 * ⁠ [9,20],
 * ⁠ [3]
 * ]
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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func getLevel(root *TreeNode, level int) []int {
	if root == nil {
		return nil
	}
	if level == 1 {
		return []int{root.Val}
	}
	var ret []int
	if level > 1 {
		ret = append(ret, getLevel(root.Left, level-1)...)
		ret = append(ret, getLevel(root.Right, level-1)...)
	}
	return ret
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	h := maxDepth(root)
	for i := h; i > 0; i-- {
		ret = append(ret, getLevel(root, i))
	}
	return ret
}

