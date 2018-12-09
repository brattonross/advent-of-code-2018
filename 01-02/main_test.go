package main

import (
	"io"
	"strings"
	"testing"
)

func TestInts(t *testing.T) {
	tests := []struct {
		reader   io.Reader
		expected []int
	}{
		{
			reader:   strings.NewReader("+1\n+1\n+1"),
			expected: []int{1, 1, 1},
		},
		{
			reader:   strings.NewReader("-1\n-2\n-3"),
			expected: []int{-1, -2, -3},
		},
	}

	for _, test := range tests {
		actual, err := ints(test.reader)
		if err != nil {
			t.Error(err)
		}
		if len(actual) != len(test.expected) {
			t.Errorf("expected length of %d, got %d", len(test.expected), len(actual))
		}
		for i, n := range actual {
			if n != test.expected[i] {
				t.Errorf("expected value %d at index %d, got %d", test.expected[i], i, n)
			}
		}
	}
}

func TestRecurringSum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, -2, 3, 1},
			expected: 2,
		},
		{
			nums:     []int{1, -1},
			expected: 0,
		},
		{
			nums:     []int{3, 3, 4, -2, -4},
			expected: 10,
		},
		{
			nums:     []int{-6, 3, 8, 5, -6},
			expected: 5,
		},
		{
			nums:     []int{7, 7, -2, -7, -4},
			expected: 14,
		},
	}

	for _, test := range tests {
		actual := recurringSum(test.nums)
		if actual != test.expected {
			t.Errorf("expected %d, got %d", test.expected, actual)
		}
	}
}
