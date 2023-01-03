//111.二叉树的最小深度
//2021-08-29 20:56:10

//给定一个二叉树，找出其最小深度。 
//
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。 
//
// 说明：叶子节点是指没有子节点的节点。 
//
// 
//
// 示例 1： 
//
// 
//输入：root = [3,9,20,null,null,15,7]
//输出：2
// 
//
// 示例 2： 
//
// 
//输入：root = [2,null,3,null,4,null,5,null,6]
//输出：5
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数的范围在 [0, 10⁵] 内 
// -1000 <= Node.val <= 1000 
// 
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 570 👎 0


package leetcode

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	curFloor := []*TreeNode{root}

	for {
		depth++
		nextFloor := make([]*TreeNode,0)
		for _, node := range curFloor {
			if node.Left == nil && node.Right == nil {
				return depth
			} else {
				if node.Left != nil {
					nextFloor = append(nextFloor, node.Left)
				}
				if node.Right != nil {
					nextFloor = append(nextFloor, node.Right)
				}
			}
		}
		curFloor = nextFloor
	}
}
//leetcode submit region end(Prohibit modification and deletion)
