package main

import "testing"

// function to test cleanInput() function
func TestCleanInput(t *testing.T) {
	// test cases with input (string) and expected output (slice of strings)
	cases := []struct{
		input string
		expected []string
	}{
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},

		{
			input: "  this   is  a test    ",
			expected: []string{"this", "is", "a", "test"},
		},

		{
			input: "1 2 3",
			expected: []string{"1", "2", "3"},
		},
	}

	// iterate over each test case
	for _, c := range cases {
		actual := cleanInput(c.input)
		// fail the test and output an error if the output is wrong length
		if len(actual) != len(c.expected) {
			t.Errorf("\033[31mFAIL\033[0m\nWant - %s\nGot  - %s", c.expected, actual)
		}
		// check each word in output for errors
		for i := range actual {
			word := actual[i]
			if word != c.expected[i] {
				t.Errorf("\033[31mFAIL\033[0m\nWant - %s\nGot  - %s", c.expected, actual)
			}
		}
	}
}
