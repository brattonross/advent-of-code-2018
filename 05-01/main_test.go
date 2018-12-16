package main

import "testing"

func TestReact(t *testing.T) {
	tests := []struct {
		units    []unit
		expected int
	}{
		{
			units: []unit{
				unit{'a', lower, false},
				unit{'a', upper, false},
			},
			expected: 0,
		},
		{
			units: []unit{
				unit{'a', lower, false},
				unit{'b', lower, false},
				unit{'b', upper, false},
				unit{'a', upper, false},
			},
			expected: 0,
		},
	}

	for _, test := range tests {
		actual := react(test.units)
		if actual != test.expected {
			t.Errorf("expected %d remaining units, got %d", test.expected, actual)
		}
	}
}
