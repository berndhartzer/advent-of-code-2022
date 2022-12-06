package aoc

import (
	"fmt"
	"testing"
	"time"
)

func tuningTroublePartOne(input string) int {
	return tuningTrouble(input, 4)
}

func tuningTroublePartTwo(input string) int {
	return tuningTrouble(input, 14)
}

func tuningTrouble(input string, size int) int {
	i := 0
	for ; i < len(input)-(size-1); i++ {
		window := input[i : i+size]
		unique := map[rune]bool{}

		for _, r := range window {
			unique[r] = true
		}

		if len(unique) == size {
			break
		}
	}

	return i + size
}

func TestDaySix(t *testing.T) {
	type testConfig struct {
		input     string
		expected  int
		logResult bool
	}

	input, err := readFileAsString("./input/06.txt")
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
	}

	runTests := func(t *testing.T, tests map[string]testConfig, fn func(string) int) {
		for name, cfg := range tests {
			cfg := cfg
			t.Run(name, func(t *testing.T) {
				start := time.Now()
				output := fn(cfg.input)
				finish := time.Since(start)
				if cfg.logResult {
					t.Log(fmt.Sprintf("\nsolution:\t%v\nelapsed time:\t%s", output, finish))
					return
				}

				if output != cfg.expected {
					t.Fatalf("Incorrect output - got: %v, want: %v", output, cfg.expected)
				}
			})
		}
	}

	t.Run("part one", func(t *testing.T) {
		tests := map[string]testConfig{
			"test_1": {
				input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
				expected: 7,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, tuningTroublePartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, tuningTroublePartTwo)
	})
}
