package twosum

import "testing"

var testCases = []struct {
	description string
	nums        []int
	target      int
	expected    []int
}{
	{
		description: "test-1",
		nums:        []int{2, 7, 11, 15},
		target:      9,
		expected:    []int{0, 1},
	}, {
		description: "test-2",
		nums:        []int{3, 2, 4},
		target:      6,
		expected:    []int{1, 2},
	}, {
		description: "test-3",
		nums:        []int{3, 3},
		target:      6,
		expected:    []int{0, 1},
	},
}

func TestTwoSum(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := TwoSum(tc.nums, tc.target)

			if !equal(tc.expected, got) {
				t.Errorf("got: %v, want: %v", got, tc.expected)
			}
		})
	}
}

func equal(expected []int, got []int) bool {
	if len(expected) != len(got) {
		return false
	}

	for i, v := range expected {
		if got[i] != v {
			return false
		}
	}

	return true
}
