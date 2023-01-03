//367.有效的完全平方数
//2021-08-23 00:52:25

//给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。 
//
// 进阶：不要 使用任何内置的库函数，如 sqrt 。 
//
// 
//
// 示例 1： 
//
// 
//输入：num = 16
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：num = 14
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// 1 <= num <= 2^31 - 1 
// 
// Related Topics 数学 二分查找 👍 235 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func isPerfectSquare(num int) bool {
	foo := mySqrtTool(num)
	return foo * foo == num
}

func mySqrtTool(x int) int {
	guess := x
	for guess*guess > x {
		guess = (guess + x/guess) / 2
	}
	return guess
}
//leetcode submit region end(Prohibit modification and deletion)
