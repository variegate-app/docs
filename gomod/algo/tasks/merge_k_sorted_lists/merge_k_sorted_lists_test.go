package merge_k_sorted_lists

import "testing"

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name     string
		args     [][]int
		expected []int
	}{
		{name: "empty", args: [][]int{}, expected: []int{}},
		{name: "single", args: [][]int{[]int{1}}, expected: []int{1}},
		{name: "full", args: [][]int{[]int{1, 4, 5}, []int{1, 3, 4}, []int{2, 6}}, expected: []int{1, 1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var lists []*ListNode
			if len(tt.args) == 0 {
				return
			}

			for _, x := range tt.args {
				list := &ListNode{}
				for _, y := range x {
					
				}
			}

			lists := &ListNode{Val: tt.args[0]}
			for _, arg := range tt.args {

			}
			args[] * ListNode
			expected * ListNode
			mergeKLists()
		})
	}
}
