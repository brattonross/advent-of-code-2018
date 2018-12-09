package main

import "testing"

func TestParseIDs(t *testing.T) {
	tests := []struct {
		ids            []string
		expectedTwice  int
		expectedThrice int
	}{
		{
			ids: []string{
				"abcdef",
				"bababc",
				"abbcde",
				"abcccd",
				"aabcdd",
				"abcdee",
				"ababab",
			},
			expectedTwice:  4,
			expectedThrice: 3,
		},
		{
			ids: []string{
				"abcdef",
				"aabbcd",
				"abcccd",
			},
			expectedTwice:  1,
			expectedThrice: 1,
		},
	}

	for _, test := range tests {
		twice, thrice := parseIDs(test.ids)
		if twice != test.expectedTwice {
			t.Errorf("expected twice to be %d, got %d", test.expectedTwice, twice)
		}
		if thrice != test.expectedThrice {
			t.Errorf("expected thrice to be %d, got %d", test.expectedThrice, thrice)
		}
	}
}
