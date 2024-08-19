// Test for leetcode problem 1. Two Sum

package leetcode

import "testing"

// TestTwoSum tests the twoSum function.
func TestTwoSum(t *testing.T) {
		test := []struct {
				nums   []int
				target int
				want   []int
		}{
				{nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
				{nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
				{nums: []int{3, 3}, target: 6, want: []int{0, 1}},
		}

		for _, tt := range test {
				got := twoSum(tt.nums, tt.target)
				if got[0] != tt.want[0] || got[1] != tt.want[1] {
						t.Errorf("twoSum(%v, %v) = %v; want %v", tt.nums, tt.target, got, tt.want)
				}
		}
}
