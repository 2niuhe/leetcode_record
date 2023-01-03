//78.子集
//2021-08-25 00:39:09

//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。 
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 
//
// 示例 2： 
//
// 
//输入：nums = [0]
//输出：[[],[0]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10 
// -10 <= nums[i] <= 10 
// nums 中的所有元素 互不相同 
// 
// Related Topics 位运算 数组 回溯 👍 1293 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	number := 1 << len(nums)
	allSubSet := make([][]int, 0, number)
	for i:= 0; i < number; i++ {
		sub := []int{}
		for j:=0; j < len(nums); j++ {
			if((i >> j) & 1) == 1 {
				sub = append(sub, nums[j])
			}
		}
		allSubSet = append(allSubSet, sub)
	}
	return allSubSet
}
//leetcode submit region end(Prohibit modification and deletion)
