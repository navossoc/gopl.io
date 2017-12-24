package main

import "testing"

func TestAnagram(t *testing.T) {
	// source: https://github.com/exercism/problem-specifications/blob/master/exercises/anagram/canonical-data.json
	testCases := []struct {
		description string
		subject     string
		candidates  []string
		expected    []bool
	}{
		{
			description: "no matches",
			subject:     "diaper",
			candidates: []string{
				"hello",
				"world",
				"zombies",
				"pants"},
			expected: []bool{false, false, false, false},
		},
		{
			description: "detects simple anagram",
			subject:     "ant",
			candidates: []string{
				"tan",
				"stand",
				"at"},
			expected: []bool{true, false, false},
		},
		{
			description: "does not detect false positives",
			subject:     "galea",
			candidates: []string{
				"eagle"},
			expected: []bool{false},
		},
		{
			description: "detects two anagrams",
			subject:     "master",
			candidates: []string{
				"stream",
				"pigeon",
				"maters"},
			expected: []bool{true, false, true},
		},
		{
			description: "does not detect anagram subsets",
			subject:     "good",
			candidates: []string{
				"dog",
				"goody"},
			expected: []bool{false, false},
		},
		{
			description: "detects anagram",
			subject:     "listen",
			candidates: []string{
				"enlists",
				"google",
				"inlets",
				"banana"},
			expected: []bool{false, false, true, false},
		},
		{
			description: "detects three anagrams",
			subject:     "allergy",
			candidates: []string{
				"gallery",
				"ballerina",
				"regally",
				"clergy",
				"largely",
				"leading"},
			expected: []bool{true, false, true, false, true, false},
		},
		{
			description: "does not detect identical words",
			subject:     "corn",
			candidates: []string{
				"corn",
				"dark",
				"Corn",
				"rank",
				"CORN",
				"cron",
				"park"},
			expected: []bool{false, false, false, false, false, true, false},
		},
		{
			description: "does not detect non-anagrams with identical checksum",
			subject:     "mass",
			candidates: []string{
				"last"},
			expected: []bool{false},
		},
		{
			description: "detects anagrams case-insensitively",
			subject:     "Orchestra",
			candidates: []string{
				"cashregister",
				"Carthorse",
				"radishes"},
			expected: []bool{false, true, false},
		},
		{
			description: "detects anagrams using case-insensitive subject",
			subject:     "Orchestra",
			candidates: []string{
				"cashregister",
				"carthorse",
				"radishes"},
			expected: []bool{false, true, false},
		},
		{
			description: "detects anagrams using case-insensitive possible matches",
			subject:     "orchestra",
			candidates: []string{
				"cashregister",
				"Carthorse",
				"radishes"},
			expected: []bool{false, true, false},
		},
		{
			description: "does not detect a word as its own anagram",
			subject:     "banana",
			candidates: []string{
				"Banana"},
			expected: []bool{false},
		},
		{
			description: "does not detect a anagram if the original word is repeated",
			subject:     "go",
			candidates: []string{
				"go Go GO"},
			expected: []bool{false},
		},
		{
			description: "anagrams must use all letters exactly once",
			subject:     "tapper",
			candidates: []string{
				"patter"},
			expected: []bool{false},
		},
		{
			description: "capital word is not own anagram",
			subject:     "BANANA",
			candidates: []string{
				"Banana"},
			expected: []bool{false},
		},
	}
	for _, tC := range testCases {

		for i, candidate := range tC.candidates {
			t.Run(tC.subject+" - "+candidate, func(t *testing.T) {

				// just a sanity check
				if len(tC.candidates) != len(tC.expected) {
					t.Fatal("something is wrong with the test data")
				}

				if anagram(tC.subject, candidate) != tC.expected[i] {
					t.Fail()
				}

			})
		}
	}
}
