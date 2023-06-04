package validparentheses

import "testing"

var testCases = []struct {
	description string
	input       string
	expect      bool
}{
	{
		description: "test1",
		input:       "()",
		expect:      true,
	},
	{
		description: "test2",
		input:       "()[]{}",
		expect:      true,
	},
	{
		description: "test3",
		input:       "(]",
		expect:      false,
	},
	{
		description: "test4-complex latex expression",
		input:       "\\left(\\begin{array}{cc} \\frac{1}{3} & x\\\\ \\mathrm{e}^{x} &... x^2 \\end{array}\\right)",
		expect:      true,
	},
}

func TestIsValid(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := IsValid(tc.input)

			if tc.expect != got {
				t.Errorf("want: %t, got: %t", tc.expect, got)
			}
		})
	}
}
