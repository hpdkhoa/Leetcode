package _16_100_SameTree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isEqualNode(p *TreeNode) bool {
	if p == nil {
		return false
	}
	return true
}

func processNode(p *TreeNode, q *TreeNode) bool {
	if isEqualNode(p) == false && isEqualNode(q) == false {
		return true
	}
	if isEqualNode(p) == true && isEqualNode(q) == true {
		if p.Val == q.Val {
			resLeft := processNode(p.Left, q.Left)
			resRight := processNode(p.Right, q.Right)
			return resLeft && resRight
		}
	}
	return false
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return processNode(p,q)
}