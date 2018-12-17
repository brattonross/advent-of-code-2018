package main

import "testing"

func TestReact(t *testing.T) {
	tests := []struct {
		a, b     unit
		expected bool
	}{
		{
			a:        unit{'a', lower},
			b:        unit{'a', upper},
			expected: true,
		},
		{
			a:        unit{'a', lower},
			b:        unit{'a', lower},
			expected: false,
		},
		{
			a:        unit{'a', upper},
			b:        unit{'a', upper},
			expected: false,
		},
		{
			a:        unit{'a', lower},
			b:        unit{'b', lower},
			expected: false,
		},
		{
			a:        unit{'a', lower},
			b:        unit{'b', upper},
			expected: false,
		},
		{
			a:        unit{'a', upper},
			b:        unit{'b', upper},
			expected: false,
		},
	}

	for _, test := range tests {
		actual := react(test.a, test.b)
		if actual != test.expected {
			t.Errorf("expected %v, got %v", test.expected, actual)
		}
	}
}

func TestRemoveUnits(t *testing.T) {
	tests := []struct {
		units    []unit
		expected []unit
		index    int
	}{
		{
			units:    []unit{unit{}, unit{}},
			expected: []unit{},
			index:    0,
		},
		{
			units:    []unit{unit{utype: 1}, unit{utype: 2}, unit{utype: 3}},
			expected: []unit{unit{utype: 3}},
			index:    0,
		},
		{
			units:    []unit{unit{utype: 1}, unit{utype: 2}, unit{utype: 3}, unit{utype: 4}},
			expected: []unit{unit{utype: 1}, unit{utype: 4}},
			index:    1,
		},
	}

	for _, test := range tests {
		actual := removeUnits(test.units, test.index)

		if len(actual) != len(test.expected) {
			t.Fatalf("expected %d elements, got %d", len(test.expected), len(actual))
		}
		for i, u := range actual {
			if u != test.expected[i] {
				t.Errorf("expected %v, got %v", test.expected[i], u)
			}
		}
	}
}
