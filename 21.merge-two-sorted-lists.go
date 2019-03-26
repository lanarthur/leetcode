import (
	"sort"
)

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (45.76%)
 * Total Accepted:    520.4K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new list. The new list
 * should be made by splicing together the nodes of the first two lists.
 *
 * Example:
 *
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func Range(m *ListNode) []int {
	var arr []int
	if m.Next == nil {
		return []int{m.Val}
	}
	t := *m
	for {
		arr = append(arr, t.Val)
		if t.Next == nil {
			break
		}
		t = *t.Next
	}
	return arr
}

func makeList(arr []int) *ListNode {
	if len(arr) == 1 {
		return &ListNode{Val: arr[0]}
	}
	return &ListNode{Val: arr[0], Next: makeList(arr[1:])}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	arr := append(Range(l1), Range(l2)...)
	sort.Ints(arr)
	return makeList(arr)
}

