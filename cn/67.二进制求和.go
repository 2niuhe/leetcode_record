//67.二进制求和
//2021-08-25 01:08:36

//给你两个二进制字符串，返回它们的和（用二进制表示）。 
//
// 输入为 非空 字符串且只包含数字 1 和 0。 
//
// 
//
// 示例 1: 
//
// 输入: a = "11", b = "1"
//输出: "100" 
//
// 示例 2: 
//
// 输入: a = "1010", b = "1011"
//输出: "10101" 
//
// 
//
// 提示： 
//
// 
// 每个字符串仅由字符 '0' 或 '1' 组成。 
// 1 <= a.length, b.length <= 10^4 
// 字符串如果不是 "0" ，就都不含前导零。 
// 
// Related Topics 位运算 数学 字符串 模拟 👍 659 👎 0


package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {
	result := []byte{}
	carry := 0
	i := len(a) - 1
	j := len(b) - 1
	for carry != 0 || i >= 0 || j >= 0 {
		total := carry
		if i >= 0 {
			total += myAtoi(a[i])
			i--
		}
		if j >= 0 {
			total += myAtoi(b[j])
			j--
		}
		result = append(result, myItoa(total%2))
		carry = total / 2
	}
	return string(reverseBytes(result))
}

func reverseBytes(foo []byte) []byte {
	bar := make([]byte, 0, len(foo))
	for i := len(foo) - 1; i >= 0; i-- {
		bar = append(bar, foo[i])
	}
	return bar
}

func myAtoi(foo byte) int {
	if foo == '0' {
		return 0
	} else {
		return 1
	}
}

func myItoa(foo int) byte {
	if foo == 0 {
		return '0'
	} else {
		return '1'
	}
}
//leetcode submit region end(Prohibit modification and deletion)
