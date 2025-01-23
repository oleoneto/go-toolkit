package helpers_test

import (
	"testing"

	"github.com/oleoneto/go-toolkit/helpers"
	"github.com/stretchr/testify/assert"
)

func TestQuicksort(t *testing.T) {
	tests := []struct {
		name        string
		lhsFunc     func(int, int, int) bool
		input       []int
		expectation []int
	}{
		{
			name:        "1234a",
			lhsFunc:     func(idx int, el, pivot int) bool { return el < pivot },
			input:       []int{1, 2, 3, 4},
			expectation: []int{1, 2, 3, 4},
		},
		{
			name:        "1234b",
			lhsFunc:     func(idx int, el, pivot int) bool { return el < pivot },
			input:       []int{2, 3, 4, 1},
			expectation: []int{1, 2, 3, 4},
		},
		{
			name:        "1234c",
			lhsFunc:     func(idx int, el, pivot int) bool { return el < pivot },
			input:       []int{2, 1, 3, 4},
			expectation: []int{1, 2, 3, 4},
		},
		{
			name:        "1234d",
			lhsFunc:     func(idx int, el, pivot int) bool { return el < pivot },
			input:       []int{4, 3, 2, 1},
			expectation: []int{1, 2, 3, 4},
		},
		{
			name:        "1234e",
			lhsFunc:     func(idx int, el, pivot int) bool { return el > pivot },
			input:       []int{3, 2, 1, 4},
			expectation: []int{4, 3, 2, 1},
		},
		{
			name:        "repeated tokens - 1",
			lhsFunc:     func(idx int, el, pivot int) bool { return el < pivot },
			input:       []int{1, 9, 4, 2, 3, 5, 4},
			expectation: []int{1, 2, 3, 4, 4, 5, 9},
		},
		{
			name:        "repeated tokens - 2",
			lhsFunc:     func(idx int, el, pivot int) bool { return el < pivot },
			input:       []int{1, 9, 4, 2, 9, 3, 5, 4},
			expectation: []int{1, 2, 3, 4, 4, 5, 9, 9},
		},
		{
			name:        "repeated tokens - 3",
			lhsFunc:     func(idx int, el, pivot int) bool { return el > pivot },
			input:       []int{1, 9, 4, 2, 9, 3, 5, 4},
			expectation: []int{9, 9, 5, 4, 4, 3, 2, 1},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectation, helpers.QuickSort(test.input, test.lhsFunc))
	}
}
