package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpackString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		err      error
	}{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", "", errors.New("(некорректная строка)")},
		{"", "", nil},
	}

	for _, tc := range testCases {
		result, err := unpackString(tc.input)

		assert.Equal(t, tc.expected, result, "ввод: %s", tc.input)

		if tc.err != nil {
			assert.EqualError(t, err, tc.err.Error(), "ввод: %s", tc.input)
		} else {
			assert.NoError(t, err, "ввод: %s", tc.input)
		}
	}
}
