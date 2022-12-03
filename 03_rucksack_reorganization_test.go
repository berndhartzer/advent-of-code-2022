package aoc

import (
	"fmt"
	"testing"
	"time"
)

func rucksackReorganizationPartOne(contents []string) int {
	dups := []byte{}

	for _, line := range contents {
		firstItems := map[byte]bool{}
		secondItems := map[byte]bool{}
		half := len(line) / 2

		for i := 0; i < half; i++ {
			firstItems[line[i]] = true
			secondItems[line[i+half]] = true
		}

		for b := range firstItems {
			_, ok := secondItems[b]
			if ok {
				dups = append(dups, b)
				break
			}
		}
	}

	total := 0
	// Loop through bytes which represent the decimal notation of the
	// character as per the ASCII table, then deduct either 96 or 38 from
	// that value to convert to the scoring value specified in the puzzle
	// (a - z is 1 - 26, etc.)
	// e.g. character p = byte(112), 112 - 96 = 16
	for _, v := range dups {
		if v > 90 {
			total += (int(v) - 96)
		} else {
			total += (int(v) - 38)
		}
	}

	return total
}

func rucksackReorganizationPartTwo(contents []string) int {
	dups := []byte{}

	groups := [3]map[byte]bool{}
	groups[0] = map[byte]bool{}
	groups[1] = map[byte]bool{}
	groups[2] = map[byte]bool{}

	finaliseGroup := func() {
		for b := range groups[0] {
			_, ok := groups[1][b]
			if !ok {
				continue
			}
			_, ok = groups[2][b]
			if !ok {
				continue
			}
			dups = append(dups, b)
			break
		}

		groups[0] = map[byte]bool{}
		groups[1] = map[byte]bool{}
		groups[2] = map[byte]bool{}
	}

	looper := 0
	for _, line := range contents {
		if looper > 2 {
			finaliseGroup()
			looper = 0
		}

		for i := 0; i < len(line); i++ {
			groups[looper][line[i]] = true
		}

		looper++
	}
	finaliseGroup()

	total := 0
	// Loop through bytes which represent the decimal notation of the
	// character as per the ASCII table, then deduct either 96 or 38 from
	// that value to convert to the scoring value specified in the puzzle
	// (a - z is 1 - 26, etc.)
	// e.g. character p = byte(112), 112 - 96 = 16
	for _, v := range dups {
		if v > 90 {
			total += (int(v) - 96)
		} else {
			total += (int(v) - 38)
		}
	}

	return total
}

func TestDayThree(t *testing.T) {
	type testConfig struct {
		input     []string
		expected  int
		logResult bool
	}

	input, err := readFileAsStringSlice("./input/03.txt")
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
			"test_1": {
				input: []string{
					"vJrwpWtwJgWrhcsFMMfFFhFp",
					"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
					"PmmdzqPrVvPwwTWBwg",
					"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
					"ttgJtRGJQctTZtZT",
					"CrZsJsPPZsGzwwsLwLmpwMDw",
				},
				expected: 157,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, rucksackReorganizationPartOne)
	})

	t.Run("part two", func(t *testing.T) {
		tests := map[string]testConfig{
			"test_1": {
				input: []string{
					"vJrwpWtwJgWrhcsFMMfFFhFp",
					"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
					"PmmdzqPrVvPwwTWBwg",
					"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
					"ttgJtRGJQctTZtZT",
					"CrZsJsPPZsGzwwsLwLmpwMDw",
				},
				expected: 70,
			},
			"solution": {
				input:     input,
				logResult: true,
			},
		}

		runTests(t, tests, rucksackReorganizationPartTwo)
	})
}
