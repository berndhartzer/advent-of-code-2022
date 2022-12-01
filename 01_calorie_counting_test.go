package aoc

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
	"time"
)

func calorieCountingPartOne(calories []string) int {
	highest := 0
	curr := 0

	for _, strCalorie := range calories {
		if strCalorie == "" {
			if curr > highest {
				highest = curr
			}

			curr = 0
			continue
		}

		calories, err := strconv.Atoi(strCalorie)
		if err != nil {
			panic(fmt.Sprintf("error converting to int: %s", err.Error()))
		}

		curr += calories
	}

	return highest
}

func calorieCountingPartTwo(calories []string) int {
	counts := []int{}
	curr := 0

	for _, strCalorie := range calories {
		if strCalorie == "" {
			counts = append(counts, curr)
			curr = 0
			continue
		}

		calories, err := strconv.Atoi(strCalorie)
		if err != nil {
			panic(fmt.Sprintf("error converting to int: %s", err.Error()))
		}

		curr += calories
	}

	sort.Ints(counts)

	return counts[len(counts)-1] + counts[len(counts)-2] + counts[len(counts)-3]
}

func TestDayOne(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/01.txt")
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

		runTests(t, tests, calorieCountingPartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, calorieCountingPartTwo)
	})
}
