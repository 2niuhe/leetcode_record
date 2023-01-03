//349.两个数组的交集
//2021-09-03 00:58:51

//给定两个数组，编写一个函数来计算它们的交集。 
//
// 
//
// 示例 1： 
//
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2]
// 
//
// 示例 2： 
//
// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[9,4] 
//
// 
//
// 说明： 
//
// 
// 输出结果中的每个元素一定是唯一的。 
// 我们可以不考虑输出结果的顺序。 
// 
// Related Topics 数组 哈希表 双指针 二分查找 排序 👍 410 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 <=0 || len2 <=0 {
		return []int{}
	}
	if len1 > len2 {
		nums2, nums1 = nums1, nums2
	}

	hashRecord := make(map[int]bool,0)
	ans := make(map[int]bool,0)
	ansL := make([]int,0)
	for _,num := range nums1 {
		hashRecord[num] = false
	}

	for _,num := range nums2 {
		_,ok := hashRecord[num]
		if ok {
			ans[num] = false
		}
	}

	for k,_ := range ans {
		ansL = append(ansL,k)
	}
	return ansL
}
//leetcode submit region end(Prohibit modification and deletion)
