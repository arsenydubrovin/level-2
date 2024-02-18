package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterUnique(t *testing.T) {
	testCases := []struct {
		input  []string
		result []string
	}{
		{[]string{"a"}, []string{"a"}},
		{[]string{"a", "b"}, []string{"a", "b"}},
		{[]string{"a", "a"}, []string{"a"}},
	}

	for _, tc := range testCases {
		result := filterUnique(tc.input)

		assert.Equal(t, tc.result, result, "%v", tc.input)
	}
}
