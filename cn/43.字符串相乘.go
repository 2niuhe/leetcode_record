//43.字符串相乘
//2021-08-26 22:56:16

//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。 
//
// 示例 1: 
//
// 输入: num1 = "2", num2 = "3"
//输出: "6" 
//
// 示例 2: 
//
// 输入: num1 = "123", num2 = "456"
//输出: "56088" 
//
// 说明： 
//
// 
// num1 和 num2 的长度小于110。 
// num1 和 num2 只包含数字 0-9。 
// num1 和 num2 均不以零开头，除非是数字 0 本身。 
// 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。 
// 
// Related Topics 数学 字符串 模拟 👍 706 👎 0
// Karatsuba快速相乘算法

package leetcode

import "strings"

//leetcode submit region begin(Prohibit modification and deletion)
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}

	num1TailZeroNum := 0
	num2TailZeroNum := 0
	for len(num1)-1-num1TailZeroNum >= 0 && num1[len(num1)-1-num1TailZeroNum] == '0' {
		num1TailZeroNum++
	}
	for len(num2)-1-num2TailZeroNum >= 0 && num2[len(num2)-1-num2TailZeroNum] == '0' {
		num2TailZeroNum++
	}

	num1 = num1[0 : len(num1)-num1TailZeroNum]
	num2 = num2[0 : len(num2)-num2TailZeroNum]

	ans := make([]string, 0)

	for i := len(num1) - 1; i >= 0; i-- {
		ans = append(ans, multiplyOne(num2, num1[i], len(num1)-i-1))
	}

	result := Reduce(ans, add)
	result += strings.Repeat("0", num1TailZeroNum+num2TailZeroNum)
	return result

}

func Reduce(arr []string, fn func(s1 string, s2 string) string) string {
	sum := "0"
	for _, it := range arr {
		sum = fn(sum, it)
	}
	return sum
}

func multiplyOne(nums1 string, nums2 byte, zeroNum int) string {
	carry := uint8(0)
	ans := make([]byte, 0)
	mul2 := nums2 - '0'
	if mul2 == 0 {
		return "0"
	}
	i := len(nums1) - 1
	for i >= 0 {
		foo := nums1[i] - '0'
		bar1 := foo*mul2 + carry
		carry = bar1 / 10
		ans = append(ans, bar1%10+'0')
		i--
	}

	for carry > 0 {
		ans = append(ans, carry%10+'0')
		carry = carry / 10
	}

	for k := 0; k < len(ans)/2; k++ {
		ans[k], ans[len(ans)-1-k] = ans[len(ans)-1-k], ans[k]
	}

	for zeroNum > 0 {
		ans = append(ans, '0')
		zeroNum--
	}
	return string(ans)
}

func add(num1 string, num2 string) string {
	carry := uint8(0)
	ans := make([]byte, 0)
	i, j := len(num1)-1, len(num2)-1
	for i >= 0 && j >= 0 {

		foo := num1[i] - '0'
		bar := num2[j] - '0'

		res := foo + bar + carry
		carry = res / 10
		ans = append(ans, res%10+'0')
		i--
		j--
	}

	for i >= 0 {
		foo := num1[i] - '0'
		res := foo + carry
		carry = res / 10
		ans = append(ans, res%10+'0')
		i--
	}

	for j >= 0 {
		foo := num2[j] - '0'
		res := foo + carry
		carry = res / 10
		ans = append(ans, res%10+'0')
		j--
	}

	for carry > 0 {
		ans = append(ans, carry%10+'0')
		carry = carry / 10
	}

	for k := 0; k < len(ans)/2; k++ {
		ans[k], ans[len(ans)-1-k] = ans[len(ans)-1-k], ans[k]
	}
	return string(ans)
}
//leetcode submit region end(Prohibit modification and deletion)
