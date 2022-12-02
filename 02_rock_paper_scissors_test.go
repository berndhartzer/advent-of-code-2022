package aoc

import (
	"fmt"
	"testing"
	"time"
)

func rockPaperScissorsPartOne(strategy []string) int {
	combinations := map[string]int{
		"A X": 1 + 3, // rock rock
		"A Y": 2 + 6, // rock paper
		"A Z": 3 + 0, // rock scissors
		"B X": 1 + 0, // paper rock
		"B Y": 2 + 3, // paper paper
		"B Z": 3 + 6, // paper scissors
		"C X": 1 + 6, // scissors rock
		"C Y": 2 + 0, // scissors paper
		"C Z": 3 + 3, // scissors scissors
	}

	return rockPaperScissors(strategy, combinations)
}

func rockPaperScissorsPartTwo(strategy []string) int {
	combinations := map[string]int{
		"A X": 3 + 0, // rock scissors
		"A Y": 1 + 3, // rock rock
		"A Z": 2 + 6, // rock paper
		"B X": 1 + 0, // paper rock
		"B Y": 2 + 3, // paper paper
		"B Z": 3 + 6, // paper scissors
		"C X": 2 + 0, // scissors paper
		"C Y": 3 + 3, // scissors scissors
		"C Z": 1 + 6, // scissors rock
	}

	return rockPaperScissors(strategy, combinations)
}

func rockPaperScissors(strategy []string, combinations map[string]int) int {
	total := 0

	for _, s := range strategy {
		total += combinations[s]
	}

	return total
}

func TestDayTwo(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/02.txt")
	if err != nil {
		t.Fatalf("failed to read input: %v", err)
	}

	runTests := func(t *testing.T, tests map[string]testConfig, fn func([]string) int) {
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
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, rockPaperScissorsPartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, rockPaperScissorsPartTwo)
	})
}
