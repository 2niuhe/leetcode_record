//145.二叉树的后序遍历
//2021-08-30 00:18:12

//给定一个二叉树，返回它的 后序 遍历。 
//
// 示例: 
//
// 输入: [1,null,2,3]  
//   1
//    \
//     2
//    /
//   3 
//
//输出: [3,2,1] 
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？ 
// Related Topics 栈 树 深度优先搜索 二叉树 👍 659 👎 0


package leetcode

func postorderTraversalOld(root *TreeNode) []int {
	ans := make([]int,0)

	var preOrder func(node *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		preOrder(root.Left)
		preOrder(root.Right)
		ans = append(ans, root.Val)
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
func postorderTraversal(root *TreeNode) []int {
	ans := make([]int,0)
	stack := make([]*TreeNode,0)

	node := root
	var prev *TreeNode

	for node != nil || len(stack)>0 {
		for node!=nil {
			//push left node
			stack = append(stack, node)
			node = node.Left
		}
		//pop out
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right == nil || node.Right == prev {
			ans = append(ans, node.Val)
			prev = node
			node = nil
		} else {
			stack = append(stack, node)
			node = node.Right
		}

	}


	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
