//104.二叉树的最大深度
//2021-08-29 22:46:09

//给定一个二叉树，找出其最大深度。 
//
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。 
//
// 说明: 叶子节点是指没有子节点的节点。 
//
// 示例： 
//给定二叉树 [3,9,20,null,null,15,7]， 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7 
//
// 返回它的最大深度 3 。 
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 955 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	curFloor := []*TreeNode{root}
	depth := 0
	for len(curFloor) > 0 {
		depth++
		nextFloor := make([]*TreeNode,0)

		for _,node := range curFloor {
			if node.Left != nil {
				nextFloor = append(nextFloor, node.Left)
			}
			if node.Right != nil {
				nextFloor = append(nextFloor, node.Right)
			}
		}
		curFloor = nextFloor
	}
	return depth
}
//leetcode submit region end(Prohibit modification and deletion)
