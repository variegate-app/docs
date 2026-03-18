package merge_k_sorted_lists

//	You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
//	Merge all the linked-lists into one sorted linked-list and return it.
//
//	Example 1:
//
//	Input: lists = [[1,4,5],[1,3,4],[2,6]]
//	Output: [1,1,2,3,4,4,5,6]
//	Explanation: The linked-lists are:
//	[
//		1->4->5,
//		1->3->4,
//		2->6
//	]
//
//	merging them into one sorted linked list:
//		1->1->2->3->4->4->5->6
//
//	Example 2:
//
//	Input: lists = []
//	Output: []
//
//	Example 3:
//
//	Input: lists = [[]]
//	Output: []
//	Constraints:
//
//	k == lists.length
//	0 <= k <= 104
//	0 <= lists[i].length <= 500
//	-104 <= lists[i][j] <= 104
//	lists[i] is sorted in ascending order.
//	The sum of lists[i].length will not exceed 104.

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	vals := make([]int, len(lists))
	pos := 0
	for _, node := range lists {
		if node == nil {
			continue
		}
		value := node.Val
		if vals[pos] <= value {
			vals = append(vals, value)
			pos++
		} else {
			for jump := pos - 1; jump >= 0; jump-- {
				if vals[jump] <= value {
					newArr := vals[:jump]
					newArr = append(newArr, value)
					newArr = append(newArr, vals[jump+1:]...)
					vals = newArr
				}
			}
		}

		if node.Next != nil {
			node = node.Next
		} else {
			node = nil
		}
	}

	return lists[0]
}
