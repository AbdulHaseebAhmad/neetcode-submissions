/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	current := root

	if current == nil {
		return nil 
	}

	newLeft := current.Right
	newRight := current.Left
	current.Left = newLeft
	current.Right = newRight

	invertTree(newLeft)
	invertTree(newRight)
	return root
}
