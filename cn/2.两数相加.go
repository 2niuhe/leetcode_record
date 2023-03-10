//2.两数相加
//2021-08-25 22:21:42

//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。 
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。 
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。 
//
// 
//
// 示例 1： 
//
// 
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
// 
//
// 示例 2： 
//
// 
//输入：l1 = [0], l2 = [0]
//输出：[0]
// 
//
// 示例 3： 
//
// 
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
// 
//
// 
//
// 提示： 
//
// 
// 每个链表中的节点数在范围 [1, 100] 内 
// 0 <= Node.val <= 9 
// 题目数据保证列表表示的数字不含前导零 
// 
// Related Topics 递归 链表 数学 👍 6632 👎 0


package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1,l2
	carry :=0
	result := &ListNode{}
	resultP := result
	// step one
	for p1 != nil && p2 != nil{
		foo := p1.Val + p2.Val + carry
		bar := &ListNode{foo % 10, nil}
		carry = foo / 10
		resultP.Next = bar
		resultP = bar
		p1 = p1.Next
		p2 = p2.Next
	}

	for p1 != nil {
		foo := p1.Val + carry
		bar := &ListNode{foo % 10, nil}
		carry = foo / 10
		resultP.Next = bar
		resultP = bar
		p1 = p1.Next
	}

	for p2 != nil {
		foo := p2.Val + carry
		bar := &ListNode{foo % 10, nil}
		carry = foo / 10
		resultP.Next = bar
		resultP = bar
		p2 = p2.Next
	}

	for carry != 0 {
		bar := &ListNode{carry % 10, nil}
		carry = carry / 10
		resultP.Next = bar
		resultP = bar
	}

	return result.Next
}
//leetcode submit region end(Prohibit modification and deletion)
