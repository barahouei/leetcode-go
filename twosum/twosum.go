// Package twosum shows sum of which numbers is equal to the target.
package twosum

// TwoSum returns a slice of indices that sum of them matches the target.
func TwoSum(nums []int, target int) []int {
	res := []int{}

	m := make(map[int]int)

	for i, num := range nums {
		if _, ok := m[num]; ok {
			res = append(res, m[num], i)
		}

		m[target-num] = i
	}

	return res
}
