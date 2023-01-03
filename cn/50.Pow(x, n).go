//50.Pow(x, n)
//2021-08-26 21:56:02

//实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，x⁴
//
// Related Topics 递归 数学 👍 721 👎 0


package leetcode



//leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {
	if x == 0.0 {
		return 0
	}
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}
	if negative {
		return 1.0 / foo(x,n)
	}

	return foo(x, n)
}

func foo(x float64, n int) float64{
	if n == 0 {
		return 1.0
	}
	carry := 1.0
	if n % 2 == 1 {
		carry *= x
	}
	bar := foo(x, n/2)
	return bar * bar * carry
}
//leetcode submit region end(Prohibit modification and deletion)
