package cmd

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitLines(t *testing.T) {
	testCases := []struct {
		input    []string
		expected [][]string
	}{
		{
			[]string{"aaa", "bbb"},
			[][]string{{"aaa"}, {"bbb"}},
		},
		{
			[]string{"aaa bbb", "ccc ddd"},
			[][]string{{"aaa", "bbb"}, {"ccc", "ddd"}},
		},
		{
			[]string{" "},
			[][]string{{}},
		},
		{
			[]string{""},
			[][]string{{}},
		},
	}

	for _, tc := range testCases {
		result := splitLines(tc.input)
		assert.Equal(t, tc.expected, result, "%v", tc.input)
	}
}

func TestCompareStrings(t *testing.T) {
	testCases := []struct {
		str1     string
		str2     string
		byNumber bool
		expected bool
	}{
		{"aa", "ab", false, true},
		{"1", "a", false, true},
		{"10", "2", false, true},
		{"1", "10", false, true},
		{"10", "1a", false, true},
		{"", "a", false, true},
		{"", "", false, false},
		{"a", "a", false, false},

		{"1", "a", true, true},
		{"10", "2", true, false},
		{"1", "10", true, true},
		{"2", "100", true, true},
		{"10", "1a", true, false},
		{"1", "", true, false},
		{"0", "1", true, true},

		{"a", "b", true, true},
		{"", "a", true, true},
	}

	for _, tc := range testCases {
		result := compareStrings(tc.str1, tc.str2, tc.byNumber)
		assert.Equal(t, tc.expected, result, "%q < %q? by number: %t", tc.str1, tc.str2, tc.byNumber)
	}
}

func TestExtractDigitalPrefix(t *testing.T) {
	testCases := []struct {
		input     string
		number    float64
		hasNumber bool
	}{
		{"1", 1, true},
		{"1.2", 1.2, true},
		{"1.2a", 1.2, true},
		{"a", math.Inf(1), false},
		{"-1", -1, true},
		{"-1a", -1, true},
		{"-1.2a", -1.2, true},
		{".2a", .2, true},
		{"-.2a", -.2, true},
		{"", math.Inf(-1), false},
	}

	for _, tc := range testCases {
		number, hasNumber := extractDigitalPrefix(tc.input)

		assert.Equal(t, tc.number, number, "%q", tc.input)
		assert.Equal(t, tc.hasNumber, hasNumber, "%q", tc.input)
	}
}

func TestCompareLines(t *testing.T) {
	testCases := []struct {
		line1    []string
		line2    []string
		byNumber bool
		col      int
		expected bool
	}{
		{[]string{"a", "b"}, []string{"b", "a"}, false, 1, true},
		{[]string{"a", "b"}, []string{"b", "a"}, false, 2, false},
		{[]string{"a", "b"}, []string{"b"}, false, 2, true},
		{[]string{"a", "b"}, []string{""}, false, 2, false},

		{[]string{"a", "2"}, []string{"a", "1"}, true, 2, false},
		{[]string{"a", "2"}, []string{"a"}, true, 2, false},
	}

	for _, tc := range testCases {
		result := compareLines(tc.line1, tc.line2, tc.byNumber, tc.col)
		assert.Equal(t, tc.expected, result, "%v < %v? by number: %t, by column: %d", tc.line1, tc.line2, tc.byNumber, tc.col)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		lines       []string
		uniqueLines []string
	}{
		{[]string{"a", "a"}, []string{"a"}},
		{[]string{"a", "b"}, []string{"a", "b"}},
		{[]string{"a", "b", "b"}, []string{"a", "b"}},
		{[]string{}, []string{}},
	}

	for _, tc := range testCases {
		uniqueLines := removeDuplicates(tc.lines)

		assert.Equal(t, tc.uniqueLines, uniqueLines, "%v", tc.lines)
	}
}
