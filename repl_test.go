package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},

		{
			input:    " Charmander Bulbasaur PICKACHU",
			expected: []string{"charmander", "bulbasaur", "pickachu"},
		},

		{
			input:    "wow coool WE",
			expected: []string{"wow", "coool", "we"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("actual length %d != expected length %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual %s != expected %s", word, expectedWord)
			}
		}
	}
}
