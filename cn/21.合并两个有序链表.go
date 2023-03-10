//21.合并两个有序链表
//2021-08-29 11:57:29

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
//
// 
//
// 示例 1： 
//
// 
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
// 
//
// 示例 2： 
//
// 
//输入：l1 = [], l2 = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：l1 = [], l2 = [0]
//输出：[0]
// 
//
// 
//
// 提示： 
//
// 
// 两个链表的节点数目范围是 [0, 50] 
// -100 <= Node.val <= 100 
// l1 和 l2 均按 非递减顺序 排列 
// 
// Related Topics 递归 链表 👍 1871 👎 0


package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: -999, Next: l1}
	pAns := head

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pAns.Next = &ListNode{Val: l1.Val, Next: nil}
			pAns = pAns.Next
			l1 =l1.Next
		} else {
			pAns.Next = &ListNode{Val: l2.Val, Next: nil}
			pAns = pAns.Next
			l2 = l2.Next
		}
	}

	if l1 != nil {
		pAns.Next = l1
	}

	if l2 != nil {
		pAns.Next = l2
	}
	return head.Next
}
//leetcode submit region end(Prohibit modification and deletion)
