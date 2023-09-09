package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// edge cases
	// if list1 is nil then return list2
	// if list2 is nil then return li1
	// if both list1 and list2 is nil then return nil
	// else
	// iterate over one list and have an output list
	// for item wise check and insert the one that's the right one
	result := &ListNode{}
	head := result
	for list1 != nil || list2 != nil {
		if list1 == nil {
			result.Next = list2
			return head
		}
		if list2 == nil {
			result.Next = list1
			return head
		}
		if list1.Val <= list2.Val {
			result.Val = list1.Val
			list1 = list1.Next
			result.Next = &ListNode{}
			result = result.Next
		} else {
			result.Val = list2.Val
			list2 = list2.Next
			result.Next = &ListNode{}
			result = result.Next
		}
	}
	return head
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	lists := mergeTwoLists(list1, list2)
	for lists != nil {
		fmt.Printf("%+v\n", lists)
		lists = lists.Next
	}
}
