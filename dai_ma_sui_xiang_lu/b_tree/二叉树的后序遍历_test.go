package b_tree

// https://leetcode.cn/problems/binary-tree-postorder-traversal/

// 后序遍历：左右中
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) (res []int) {
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		inorder(root.Right)
		res = append(res, root.Val)
	}
	inorder(root)

	return
}
