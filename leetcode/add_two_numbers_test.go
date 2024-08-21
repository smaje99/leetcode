// Test for leetcode problem 2. Add two numbers

package leetcode

import (
	"reflect"
	"testing"
)

func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}

	return dummy.Next
}

func listToSlice(l *ListNode) []int {
	var result []int

	for l != nil {
		result = append(result, l.Val)
		l = l.Next
	}

	return result
}

func TestAddTwoNumbers(t *testing.T) {
	test := []struct {
		l1       []int
		l2       []int
		expected []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for _, tt := range test {
		l1 := buildList(tt.l1)
		l2 := buildList(tt.l2)
		expected := tt.expected

		actual := listToSlice(addTwoNumbers(l1, l2))

		if len(actual) != len(expected) {
			t.Fatalf("expected: %v, got: %v", expected, actual)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("addTwoNumbers(%v, %v) = %v; expected %v", l1, l2, actual, expected)
		}
	}
}
