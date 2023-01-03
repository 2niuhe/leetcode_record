//229.求众数 II
//2021-08-31 23:47:49

//给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。 
//
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。 
//
// 
//
// 示例 1： 
//
// 
//输入：[3,2,3]
//输出：[3] 
//
// 示例 2： 
//
// 
//输入：nums = [1]
//输出：[1]
// 
//
// 示例 3： 
//
// 
//输入：[1,1,1,3,3,2,2,2]
//输出：[1,2] 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 5 * 10⁴ 
// -10⁹ <= nums[i] <= 10⁹ 
// 
// Related Topics 数组 哈希表 计数 排序 👍 395 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement2(nums []int) []int {
	m := make(map[int]int)
	ans := make([]int,0)
	for _,num := range nums {
		_,ok := m[num]
		if ok {
			m[num] = m[num]+1
		} else {
			m[num] = 1
		}
	}
	line := len(nums)/3
	for key, val := range m {
		if val > line {
			ans = append(ans, key)
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
