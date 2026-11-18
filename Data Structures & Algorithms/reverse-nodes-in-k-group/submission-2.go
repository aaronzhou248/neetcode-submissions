/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	current := head // 4
	tempHead := head // 4
	outputHead := head // 3
	var preNode *ListNode // 1
	var postNode *ListNode // nil

	for current != nil { // 4 != nil
		i := 1 // 2
		for i < k && current != nil { // 1 < 3 && 5 != nil
			i++
			current = current.Next
		}
		if i == k && current != nil { // 3 == 3
			postNode = current.Next // postNode = nil
			newHead := reverseList(preNode, tempHead, postNode) // newHead = 6
			if preNode == nil { // 1 == nil
				outputHead = newHead
			}
			preNode = tempHead
			tempHead = postNode
			current = postNode
			postNode = nil
		} else {
			break
		}
	}
	return outputHead
}

func reverseList(preNode *ListNode, head *ListNode, postNode *ListNode) *ListNode {
	if head == nil {
		return head
	}

	curr := head
	prev := postNode

	for curr != postNode {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	if preNode != nil {
		preNode.Next = prev
	}
	return prev
}

// ===== Call 1: reverseList(pre=nil, head=1, post=4) =====
// start: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> nil
// init:  prev=4 (post)  curr=1
//
// iter | temp | changed   | prev | curr
//  1   |  2   | 1.Next=4  |  1   |  2
//  2   |  3   | 2.Next=1  |  2   |  3
//  3   |  4   | 3.Next=2  |  3   |  4    <- curr == post, exit
//
// next: 1→4  2→1  3→2  4→5  5→6  6→nil
// list (from 3): 3 -> 2 -> 1 -> 4 -> 5 -> 6 -> nil

// ===== Call 2: reverseList(pre=1, head=4, post=nil) =====
// start: 3 -> 2 -> 1 -> 4 -> 5 -> 6 -> nil
// init:  prev=nil (post)  curr=4
//
// iter | temp | changed   | prev | curr
//  1   |  5   | 4.Next=nil|  4   |  5
//  2   |  6   | 5.Next=4  |  5   |  6
//  3   |  nil | 6.Next=5  |  6   |  nil    <- curr == post, exit
//
// if curr == nil
//   preNode(1).Next = prev (6)
// next: 1→6  2→1  3→2  4→nil  5→4  6→5
// list (from 3): 3 -> 2 -> 1 -> 6 -> 5 -> 4 -> nil