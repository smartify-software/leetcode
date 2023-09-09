package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//numbers := addTwoNumbers(&ListNode{
	//	Val:  2,
	//	Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}},
	//}, &ListNode{
	//	Val:  5,
	//	Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}},
	//})
	//for numbers != nil {
	//	println(numbers.Val)
	//	numbers = numbers.Next
	//}

	numbers1 := addTwoNumbers(&ListNode{
		Val:  9,
		Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}},
	}, &ListNode{
		Val:  9,
		Next: &ListNode{Val: 9},
	})
	for numbers1 != nil {
		println(numbers1.Val)
		numbers1 = numbers1.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	val := l2.Val
	if l1.Val+l2.Val >= 10 {
		val = (l1.Val + l2.Val) % 10
		if l1.Next != nil {
			l1.Next.Val += 1
		} else {
			l1.Next = &ListNode{Val: 1}
		}
		return &ListNode{Val: val, Next: addTwoNumbers(l1.Next, l2.Next)}
	}
	return &ListNode{Val: l1.Val + val, Next: addTwoNumbers(l1.Next, l2.Next)}
}
