//217.存在重复元素
//2021-09-03 00:23:50

//给定一个整数数组，判断是否存在重复元素。 
//
// 如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。 
//
// 
//
// 示例 1: 
//
// 
//输入: [1,2,3,1]
//输出: true 
//
// 示例 2: 
//
// 
//输入: [1,2,3,4]
//输出: false 
//
// 示例 3: 
//
// 
//输入: [1,1,1,3,3,4,3,2,4,2]
//输出: true 
// Related Topics 数组 哈希表 排序 👍 461 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func containsDuplicate(nums []int) bool {
	hashRecord := make(map[int]int,0)

	for _,num := range nums {
		_,ok := hashRecord[num]
		if ok {
			return true
		} else {
			hashRecord[num] = 0
		}
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
