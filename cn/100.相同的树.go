//100.相同的树
//2021-08-29 22:03:07

//给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。 
//
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。 
//
// 
//
// 示例 1： 
//
// 
//输入：p = [1,2,3], q = [1,2,3]
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：p = [1,2], q = [1,null,2]
//输出：false
// 
//
// 示例 3： 
//
// 
//输入：p = [1,2,1], q = [1,1,2]
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// 两棵树上的节点数目都在范围 [0, 100] 内 
// -10⁴ <= Node.val <= 10⁴ 
// 
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 674 👎 0


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
func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if !mayEqual(p,q) {
		return false
	}

	curPFloor := []*TreeNode{p}
	curQFloor := []*TreeNode{q}

	for len(curPFloor) > 0 {
		nextPFloor := make([]*TreeNode,0)
		nextQFloor := make([]*TreeNode,0)
		for i:=0;i<len(curPFloor);i++ {
			if curPFloor[i].Val == curQFloor[i].Val &&
				mayEqual(curPFloor[i].Left, curQFloor[i].Left) &&
				mayEqual(curPFloor[i].Right, curQFloor[i].Right){
				if curPFloor[i].Left != nil {
					nextPFloor = append(nextPFloor, curPFloor[i].Left)
					nextQFloor = append(nextQFloor, curQFloor[i].Left)
				}
				if curPFloor[i].Right != nil {
					nextPFloor = append(nextPFloor, curPFloor[i].Right)
					nextQFloor = append(nextQFloor, curQFloor[i].Right)
				}
			} else {
				return false
			}

		}
		curPFloor = nextPFloor
		curQFloor = nextQFloor

	}

	return true
}


func mayEqual(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil {
		return true
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
