//107.二叉树的层序遍历 II
//2021-09-04 22:51:49

//给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历） 
//
// 例如： 
//给定二叉树 [3,9,20,null,null,15,7], 
//
// 
//    3
//   / \
//  9  20
//    /  \
//   15   7
// 
//
// 返回其自底向上的层序遍历为： 
//
// 
//[
//  [15,7],
//  [9,20],
//  [3]
//]
// 
// Related Topics 树 广度优先搜索 二叉树 👍 473 👎 0


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
func levelOrderBottom(root *TreeNode) [][]int {
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
	for i:= len(ansNode)-1; i>=0; i--{
		foo := make([]int, len(ansNode[i]))
		for j:=0;j<len(ansNode[i]);j++ {
			foo[j] = ansNode[i][j].Val
		}
		ans = append(ans, foo)
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
