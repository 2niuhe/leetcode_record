//19.删除链表的倒数第 N 个结点
//2021-08-30 23:17:28

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。 
//
// 进阶：你能尝试使用一趟扫描实现吗？ 
//
// 
//
// 示例 1： 
//
// 
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
// 
//
// 示例 2： 
//
// 
//输入：head = [1], n = 1
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：head = [1,2], n = 1
//输出：[1]
// 
//
// 
//
// 提示： 
//
// 
// 链表中结点的数目为 sz 
// 1 <= sz <= 30 
// 0 <= Node.val <= 100 
// 1 <= n <= sz 
// 
// Related Topics 链表 双指针 👍 1533 👎 0


package leetcode
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first := head
	target := head

	dummyHead := &ListNode{0,head}
	targetBefore:= dummyHead

	for i:=0;i<n-1;i++{
		first = first.Next
	}

	for first.Next != nil {
		first = first.Next
		target = target.Next
		targetBefore = targetBefore.Next
	}

	targetBefore.Next = target.Next
	return dummyHead.Next
}
//leetcode submit region end(Prohibit modification and deletion)
