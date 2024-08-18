// Leetcode problem: 1. Two Sum
// URL: https://leetcode.com/problems/two-sum/

package leetcode

// TwoSum finds two numbers in the array 'nums' that sum up to the 'target' value.
// It returns the indices of the two numbers as an array. If no such numbers are found,
// it returns nil. Assumes that there is exactly one solution.
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, found := m[complement]; found {
			return []int{index, i}
		}
		m[num] = i
	}

	return nil
}
