//204.计数质数
//2021-09-03 00:07:16

//统计所有小于非负整数 n 的质数的数量。 
//
// 
//
// 示例 1： 
//
// 输入：n = 10
//输出：4
//解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
// 
//
// 示例 2： 
//
// 输入：n = 0
//输出：0
// 
//
// 示例 3： 
//
// 输入：n = 1
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// 0 <= n <= 5 * 10⁶ 
// 
// Related Topics 数组 数学 枚举 数论 👍 738 👎 0


package leetcode


//leetcode submit region begin(Prohibit modification and deletion)
func countPrimes(n int) int {
	isPrimes := make([]bool, n)
	for i,_ := range isPrimes{
		isPrimes[i] = true
	}
	cnt := 0
	for x:=2;x<n;x++ {
		if isPrimes[x] == true {
			cnt++
			for j:= x*x; j<n; j+=x {
				isPrimes[j] = false
			}
		}

	}
	return cnt
}
//leetcode submit region end(Prohibit modification and deletion)
