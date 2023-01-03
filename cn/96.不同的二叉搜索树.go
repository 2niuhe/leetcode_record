//96.不同的二叉搜索树
//2021-09-04 21:03:32

//给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。 
//
// 
//
// 示例 1： 
//
// 
//输入：n = 3
//输出：5
// 
//
// 示例 2： 
//
// 
//输入：n = 1
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 19 
// 
// Related Topics 树 二叉搜索树 数学 动态规划 二叉树 👍 1303 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func numTrees(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1,1

	for i:=2;i<=n;i++{
		for j:=1;j<=i;j++{
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}
//leetcode submit region end(Prohibit modification and deletion)
