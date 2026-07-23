package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "Squirtle",
			expected: []string{"squirtle"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf(
				"Length mismatch for input %q: expected %d words, got %d",
				c.input,
				len(c.expected),
				len(actual),
			)
			continue
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf(
					"Input %q: expected %q, got %q",
					c.input,
					c.expected[i],
					actual[i],
				)
			}
		}
	}
}
