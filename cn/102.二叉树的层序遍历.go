//102.二叉树的层序遍历
//2021-09-04 22:32:23

//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。 
//
// 
//
// 示例： 
//二叉树：[3,9,20,null,null,15,7], 
//
// 
//    3
//   / \
//  9  20
//    /  \
//   15   7
// 
//
// 返回其层序遍历结果： 
//
// 
//[
//  [3],
//  [9,20],
//  [15,7]
//]
// 
// Related Topics 树 广度优先搜索 二叉树 👍 991 👎 0


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
func levelOrder(root *TreeNode) [][]int {
	ansNode := make([][]*TreeNode,0)

	level := make([]*TreeNode,0)
	if root != nil {
		level = append(level, root)
	}
	for len(level) > 0 {
		ansNode = append(ansNode, level)
		nextLevel := make([]*TreeNode,0)
		for _,node := range level{
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		level = nextLevel
	}
	ans := make([][]int, 0, len(ansNode))
	for _, level := range ansNode {
		foo := make([]int, len(level))
		for i:=0;i<len(level);i++ {
			foo[i] = level[i].Val
		}
		ans = append(ans, foo)
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
