//58.最后一个单词的长度
//2021-08-28 00:10:13

//给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中最后一个单词的长度。 
//
// 单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "Hello World"
//输出：5
// 
//
// 示例 2： 
//
// 
//输入：s = "   fly me   to   the moon  "
//输出：4
// 
//
// 示例 3： 
//
// 
//输入：s = "luffy is still joyboy"
//输出：6
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 10⁴ 
// s 仅有英文字母和空格 ' ' 组成 
// s 中至少存在一个单词 
// 
// Related Topics 字符串 👍 345 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLastWord(s string) int {
	firstIndex := len(s)-1

	for firstIndex >=0 && s[firstIndex] == ' '{
		firstIndex--
	}

	secondIndex := firstIndex

	for secondIndex >= 0 && s[secondIndex] != ' '{
		secondIndex--
	}

	if firstIndex == -1 {
		return 1
	}

	return firstIndex-secondIndex

}
//leetcode submit region end(Prohibit modification and deletion)
