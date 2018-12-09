package main

import (
	"reflect"
	"testing"
)

func TestParseClaims(t *testing.T) {
	tests := []struct {
		claims   []string
		expected map[int]map[int]int
	}{
		{
			claims: []string{"#1 @ 2,3: 4x4"},
			expected: map[int]map[int]int{
				2: map[int]int{
					3: 1,
					4: 1,
					5: 1,
					6: 1,
				},
				3: map[int]int{
					3: 1,
					4: 1,
					5: 1,
					6: 1,
				},
				4: map[int]int{
					3: 1,
					4: 1,
					5: 1,
					6: 1,
				},
				5: map[int]int{
					3: 1,
					4: 1,
					5: 1,
					6: 1,
				},
			},
		},
		{
			claims: []string{"#1 @ 2,3: 4x4", "#2 @ 2,3: 1x1"},
			expected: map[int]map[int]int{
				2: map[int]int{
					3: 2,
					4: 1,
					5: 1,
					6: 1,
				},
				3: map[int]int{
					3: 1,
					4: 1,
					5: 1,
					6: 1,
				},
				4: map[int]int{
					3: 1,
					4: 1,
					5: 1,
					6: 1,
				},
				5: map[int]int{
					3: 1,
					4: 1,
					5: 1,
					6: 1,
				},
			},
		},
		{
			claims: []string{"#1 @ 2,3: 4x4", "#2 @ 2,3: 1x1", "#3 @ 2,3: 2x2"},
			expected: map[int]map[int]int{
				2: map[int]int{
					3: 3,
					4: 2,
					5: 1,
					6: 1,
				},
				3: map[int]int{
					3: 2,
					4: 2,
					5: 1,
					6: 1,
				},
				4: map[int]int{
					3: 1,
					4: 1,
					5: 1,
					6: 1,
				},
				5: map[int]int{
					3: 1,
					4: 1,
					5: 1,
					6: 1,
				},
			},
		},
	}

	for _, test := range tests {
		actual := parseClaims(test.claims)

		if !reflect.DeepEqual(test.expected, actual) {
			t.Error("maps are not equal")
		}
	}
}

func TestCountOverlaps(t *testing.T) {
	tests := []struct {
		fabric   map[int]map[int]int
		expected int
	}{
		{
			fabric: map[int]map[int]int{
				2: map[int]int{
					3: 3,
					4: 2,
					5: 1,
					6: 1,
				},
				3: map[int]int{
					3: 2,
					4: 2,
					5: 1,
					6: 1,
				},
			},
			expected: 4,
		},
		{
			fabric: map[int]map[int]int{
				2: map[int]int{
					3: 1,
					4: 1,
					5: 1,
					6: 1,
				},
				3: map[int]int{
					3: 1,
					4: 1,
					5: 1,
					6: 1,
				},
			},
			expected: 0,
		},
	}

	for _, test := range tests {
		actual := countOverlaps(test.fabric)
		if actual != test.expected {
			t.Errorf("expected %d, got %d", test.expected, actual)
		}
	}
}
