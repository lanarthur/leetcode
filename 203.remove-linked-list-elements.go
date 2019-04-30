/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	ret := &ListNode{Val: 0, Next: head}
	index := ret
	for index.Next != nil {
		if index.Next.Val == val {
			next := index.Next
			index.Next = next.Next
		} else {
			index = index.Next
		}
	}
	return ret.Next
}

