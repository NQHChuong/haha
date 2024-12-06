package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}
func listToSlice(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{}
	result := list
	carry := 0

	for l1 != nil || l2 != nil {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + carry
		carry = sum / 10
		result.Next = &ListNode{Val: sum % 10}
		result = result.Next
	}
	if carry > 0 {
		result.Next = &ListNode{Val: carry}
	}
	return list.Next
}
func main() {

	l1 := []int{9, 9, 9, 9, 9}
	l2 := []int{9, 9, 9, 9, 9}
	list1 := sliceToList(l1)
	list2 := sliceToList(l2)
	result := AddTwoNumbers(list1, list2)
	fmt.Println("Result is: ", listToSlice(result))
}
