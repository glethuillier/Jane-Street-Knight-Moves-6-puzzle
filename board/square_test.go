package board

import (
	"testing"
)

type testCase struct {
	name   string
	input  Coordinates
	output string
}

func TestString(t *testing.T) {
	testCases := []testCase{
		{
			name:   "FileA_Rank1",
			input:  Coordinates{File: 'a', Rank: 1},
			output: "a1",
		},
		{
			name:   "FileF_Rank6",
			input:  Coordinates{File: 'f', Rank: 6},
			output: "f6",
		},
		{
			name:   "FileD_Rank5",
			input:  Coordinates{File: 'd', Rank: 5},
			output: "d5",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.input.ToString()
			expected := tc.output
			if actual != expected {
				t.Errorf("Expected: %s, but got: %s", expected, actual)
			}
		})
	}
}
