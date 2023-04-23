package twosum

func TwoSum(nums []int, target int) []int {
	res := []int{}

	for i, left := range nums {
		for j, right := range nums {
			if left+right == target && i != j {
				res = append(res, i)
			}
		}
	}

	return res
}
