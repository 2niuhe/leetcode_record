//101.对称二叉树
//2021-09-04 21:57:51

//给定一个二叉树，检查它是否是镜像对称的。 
//
// 
//
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。 
//
//     1
//   / \
//  2   2
// / \ / \
//3  4 4  3
// 
//
// 
//
// 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的: 
//
//     1
//   / \
//  2   2
//   \   \
//   3    3
// 
//
// 
//
// 进阶： 
//
// 你可以运用递归和迭代两种方法解决这个问题吗？ 
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1516 👎 0




package leetcode
//solution 1
func isSymmetric1(root *TreeNode) bool {
	level := []*TreeNode{root}

	for len(level) > 0 {
		nextLevel := make([]*TreeNode,0)
		for _,node := range level {
			if !checkMirror(level){
				return false
			}
			if node != nil {
				nextLevel = append(nextLevel, node.Left, node.Right)
			}
		}
		level = nextLevel
	}
	return true
}

func checkMirror(level []*TreeNode) bool {
	if len(level) <=1 {
		return true
	}
	l,r := 0, len(level)-1
	for l <= r {

		if level[l] == nil && level[r] != nil {
			return false
		}
		if level[l] != nil && level[r] == nil {
			return false
		}
		if level[l] != nil && level[r] != nil {
			if level[l].Val != level[r].Val {
				return false
			}
		}
		l++
		r--
	}
	return true
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
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

//leetcode submit region end(Prohibit modification and deletion)
