package _04_21_MergeTwoSortedLists


 type ListNode struct {
     Val int
     Next *ListNode
 }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := new(ListNode)
	if l1.Val < l2.Val {
		head.Val = l1.Val
		l1 = l1.Next
	} else {
		head.Val = l2.Val
		l2 = l2.Next
	}
	node := head
	for {
		if l1 != nil{
			if l2 != nil{
				if l1.Val < l2.Val {
					node.Next = l1
					l1 = l1.Next
				} else {
					node.Next = l2
					l2 = l2.Next
				}
				node = node.Next
			} else {
				node.Next = l1
				break
			}
		} else {
			if l2 != nil{
				node.Next = l2
				break
			} else {
				break
			}
		}
	}
	return head
}
