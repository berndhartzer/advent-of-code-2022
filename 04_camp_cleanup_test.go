package aoc

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func campCleanupPartOne(sections []string) int {
	overlaps := 0

	for _, line := range sections {
		split := strings.FieldsFunc(line, func(r rune) bool {
			return r == '-' || r == ','
		})

		firstLower, err := strconv.Atoi(split[0])
		if err != nil {
			panic("failed to convert string to int")
		}
		firstUpper, err := strconv.Atoi(split[1])
		if err != nil {
			panic("failed to convert string to int")
		}
		secondLower, err := strconv.Atoi(split[2])
		if err != nil {
			panic("failed to convert string to int")
		}
		secondUpper, err := strconv.Atoi(split[3])
		if err != nil {
			panic("failed to convert string to int")
		}

		if firstLower <= secondLower && firstUpper >= secondUpper {
			overlaps++
			continue
		}
		if secondLower <= firstLower && secondUpper >= firstUpper {
			overlaps++
			continue
		}
	}

	return overlaps
}

func campCleanupPartTwo(sections []string) int {
	overlaps := 0

	for _, line := range sections {
		split := strings.FieldsFunc(line, func(r rune) bool {
			return r == '-' || r == ','
		})

		firstLower, err := strconv.Atoi(split[0])
		if err != nil {
			panic("failed to convert string to int")
		}
		firstUpper, err := strconv.Atoi(split[1])
		if err != nil {
			panic("failed to convert string to int")
		}
		secondLower, err := strconv.Atoi(split[2])
		if err != nil {
			panic("failed to convert string to int")
		}
		secondUpper, err := strconv.Atoi(split[3])
		if err != nil {
			panic("failed to convert string to int")
		}

		if firstLower <= secondLower && firstUpper >= secondLower {
			overlaps++
			continue
		}
		if secondLower <= firstLower && secondUpper >= firstLower {
			overlaps++
			continue
		}
	}

	return overlaps
}

func TestDayFour(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/04.txt")
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

		runTests(t, tests, campCleanupPartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, campCleanupPartTwo)
	})
}
