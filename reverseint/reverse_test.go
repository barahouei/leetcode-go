package reverseint

import "testing"

var testCases = []struct {
	desc     string
	input    int
	expected int
}{
	{
		desc:     "test1",
		input:    123,
		expected: 321,
	},
	{
		desc:     "test2",
		input:    -123,
		expected: -321,
	},
	{
		desc:     "test3",
		input:    120,
		expected: 21,
	},
}

func TestReverse(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := Reverse(tc.input)

			if got != tc.expected {
				t.Errorf("got: %d, want: %d", got, tc.expected)
			}
		})
	}
}
