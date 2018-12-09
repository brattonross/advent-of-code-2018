package main

import "testing"

func TestCommonChars(t *testing.T) {
	tests := []struct {
		a, b, expected string
	}{
		{
			a:        "abcd",
			b:        "abed",
			expected: "abd",
		},
		{
			a:        "abab",
			b:        "bbaa",
			expected: "ba",
		},
	}

	for _, test := range tests {
		actual := commonChars(test.a, test.b)
		if actual != test.expected {
			t.Errorf("expected %s, got %s", test.expected, actual)
		}
	}
}

func TestClosestStrings(t *testing.T) {
	tests := []struct {
		strings   []string
		expectedA string
		expectedB string
	}{
		{
			strings: []string{
				"abcd",
				"abdc",
				"eeee",
			},
			expectedA: "abcd",
			expectedB: "abdc",
		},
		{
			strings: []string{
				"aaaa",
				"aaab",
				"aabb",
			},
			expectedA: "aaaa",
			expectedB: "aaab",
		},
		{
			strings: []string{
				"abab",
				"abbb",
				"abab",
			},
			expectedA: "abab",
			expectedB: "abab",
		},
	}

	for _, test := range tests {
		a, b := closestStrings(test.strings)
		if a != test.expectedA {
			t.Errorf("expected string a to be %s, got %s", test.expectedA, a)
		}
		if b != test.expectedB {
			t.Errorf("expected string b to be %s, got %s", test.expectedB, b)
		}
	}
}

func TestCommonBytes(t *testing.T) {
	tests := []struct {
		stringA  string
		stringB  string
		expected int
	}{
		{
			stringA:  "abcd",
			stringB:  "abce",
			expected: 3,
		},
		{
			stringA:  "aaaa",
			stringB:  "bbbb",
			expected: 0,
		},
		{
			stringA:  "aaaa",
			stringB:  "abab",
			expected: 2,
		},
	}

	for _, test := range tests {
		actual := commonBytes(test.stringA, test.stringB)
		if actual != test.expected {
			t.Errorf("expected %d, got %d", test.expected, actual)
		}
	}
}
