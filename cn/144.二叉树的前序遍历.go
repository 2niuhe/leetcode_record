//144.二叉树的前序遍历
//2021-08-29 22:58:19

//给你二叉树的根节点 root ，返回它节点值的 前序 遍历。 
//
// 
//
// 示例 1： 
//
// 
//输入：root = [1,null,2,3]
//输出：[1,2,3]
// 
//
// 示例 2： 
//
// 
//输入：root = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：root = [1]
//输出：[1]
// 
//
// 示例 4： 
//
// 
//输入：root = [1,2]
//输出：[1,2]
// 
//
// 示例 5： 
//
// 
//输入：root = [1,null,2]
//输出：[1,2]
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数目在范围 [0, 100] 内 
// -100 <= Node.val <= 100 
// 
//
// 
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？ 
// Related Topics 栈 树 深度优先搜索 二叉树 👍 633 👎 0


package leetcode

func preorderTraversalOld(root *TreeNode) []int {
	ans := make([]int,0)

	var preOrder func(node *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		preOrder(root.Left)
		preOrder(root.Right)
	}

	preOrder(root)
	return ans
}

//TODO:  use stack

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode,0)

	ans := make([]int,0)
	node := root

	for node!=nil || len(stack)>0 {
		for node!=nil {
			ans = append(ans, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
