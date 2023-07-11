package main

import "fmt"

/* https://leetcode.com/problems/add-two-numbers/
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) print() {
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Next: nil, Val: 0}
	sumTail := head
	tens := 0
	for l1 != nil || l2 != nil {
		sum := tens
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sumTail.Next = &ListNode{Val: sum % 10}
		tens = sum / 10
		sumTail = sumTail.Next
	}

	if tens > 0 {
		sumTail.Next = &ListNode{Val: tens}
	}

	return head.Next
}

func main() {
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}

	sum := addTwoNumbers(l2, l1)
	sum.print()
}
