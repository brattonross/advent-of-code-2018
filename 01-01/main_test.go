package main

import (
	"io"
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		reader   io.Reader
		expected int
	}{
		{
			reader:   strings.NewReader("+1\n+1\n+1"),
			expected: 3,
		},
		{
			reader:   strings.NewReader("+1\n+1\n-2"),
			expected: 0,
		},
		{
			reader:   strings.NewReader("-1\n-2\n-3"),
			expected: -6,
		},
	}

	for _, test := range tests {
		actual, err := sum(test.reader)
		if err != nil {
			t.Error(err)
		}
		if actual != test.expected {
			t.Errorf("expected %d, got %d", test.expected, actual)
		}
	}
}
