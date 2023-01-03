//94.二叉树的中序遍历
//2021-08-30 00:47:34

//给定一个二叉树的根节点 root ，返回它的 中序 遍历。 
//
// 
//
// 示例 1： 
//
// 
//输入：root = [1,null,2,3]
//输出：[1,3,2]
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
//输出：[2,1]
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
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？ 
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1080 👎 0

package leetcode

func inorderTraversalOld(root *TreeNode) []int {
	ans := make([]int,0)

	var preOrder func(node *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		preOrder(root.Left)
		ans = append(ans, root.Val)
		preOrder(root.Right)

	}

	preOrder(root)
	return ans
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	ans := make([]int,0)
	stack := make([]*TreeNode,0)

	node := root

	for node != nil || len(stack)>0 {
		for node!=nil {
			//push left node
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		node = node.Right

	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
