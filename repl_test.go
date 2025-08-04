package main

import(
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "  heLlo  world  ",
		expected: []string{"hello", "world"},
	},
	{
		input:    "woRld",
		expected: []string{"world"},
	},
}

for _, c := range cases {
	actual := cleanInput(c.input)
	if len(actual) != len(c.expected) {
		t.Errorf("Length not equal. Actual: %d Expected: %d", len(actual), len(c.expected))
	}
	// Check the length of the actual slice against the expected slice
	// if they don't match, use t.Errorf to print an error message
	// and fail the test

	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		// Check each word in the slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if word != expectedWord {
			t.Errorf("Words not matching: %s, %s", word, expectedWord)
		}
	}
}
}