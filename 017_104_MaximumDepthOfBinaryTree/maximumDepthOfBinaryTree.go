package _17_104_MaximumDepthOfBinaryTree

  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := 1
	right := 1
	if root.Left != nil {
		left += maxDepth(root.Left)
	}
	if root.Right != nil {
		right += maxDepth(root.Right)
	}
	if right > left {
		return right
	}
	return left
}