//387.字符串中的第一个唯一字符
//2021-09-04 00:06:46

//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。 
//
// 
//
// 示例： 
//
// s = "leetcode"
//返回 0
//
//s = "loveleetcode"
//返回 2
// 
//
// 
//
// 提示：你可以假定该字符串只包含小写字母。 
// Related Topics 队列 哈希表 字符串 计数 👍 431 👎 0


package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func firstUniqChar(s string) int {
	hashRecord := make(map[byte]int, 0)
	for k, _ := range s {
		_, ok := hashRecord[s[k]]
		if ok {
			hashRecord[s[k]] = math.MaxInt32
		} else {
			hashRecord[s[k]] = k
		}
	}
	minIndex := math.MaxInt32

	for _,v := range hashRecord {
		if v < minIndex {
			minIndex = v
		}
	}

	if minIndex < math.MaxInt32 {
		return minIndex
	} else {
		return -1
	}
}
//leetcode submit region end(Prohibit modification and deletion)
