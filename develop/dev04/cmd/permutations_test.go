package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPermutations(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{"", []string{}},
		{"a", []string{"a"}},
		{"ab", []string{"ab", "ba"}},
		{"abc", []string{"abc", "acb", "bac", "bca", "cba", "cab"}},
	}

	for _, tc := range testCases {
		permutations := getPermutations(tc.input)
		result := permutations
		assert.Equal(t, tc.expected, result, "%v", tc.input)
	}
}
