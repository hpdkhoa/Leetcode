package _08_83_RemoveDuplicatedfromSortedList

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	for node := head ; ; {
		if node.Next == nil {
			return head
		} else {
			if node.Next == nil {
				return head
			}
			if node.Val == node.Next.Val {
				node.Next.Val = 0
				node.Next = node.Next.Next
			} else {
				node = node.Next
			}
		}
	}
}