package cmd

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadLines(t *testing.T) {
	testCases := []struct {
		source   string
		expected map[string]struct{}
	}{
		{"a", map[string]struct{}{"a": {}}},
		{"a\nb", map[string]struct{}{"a": {}, "b": {}}},
		{"a\nb\n", map[string]struct{}{"a": {}, "b": {}}},
		{"\n", map[string]struct{}{"": {}}},
		{"", map[string]struct{}{}},
	}

	for _, tc := range testCases {
		keysCnt := len(tc.expected)
		dict := make(map[string]struct{}, keysCnt)
		r := strings.NewReader(tc.source)
		_ = readLines(r, dict)

		assert.Equal(t, tc.expected, dict, "%q", tc.source)
	}
}

func TestCountLines(t *testing.T) {
	testCases := []struct {
		source string
		count  int
	}{
		{"a", 0},
		{"a\n", 1},
		{"a\nb", 1},
		{"a\nb\n", 2},
		{"\n", 1},
		{"", 0},
	}

	for _, tc := range testCases {
		r := strings.NewReader(tc.source)
		count, _ := countLines(r)

		assert.Equal(t, tc.count, count, "%q", tc.source)
	}
}
